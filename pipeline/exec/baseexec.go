package exec

import (
	"context"
	"errors"
	"fmt"

	"github.com/streamingfast/substreams/reqctx"
	"github.com/streamingfast/substreams/storage/execout"
	"github.com/streamingfast/substreams/storage/index"
	"github.com/streamingfast/substreams/wasm"
	ttrace "go.opentelemetry.io/otel/trace"
)

var ErrWasmDeterministicExec = errors.New("wasm execution failed deterministically")

type BaseExecutor struct {
	ctx context.Context

	moduleName    string
	initialBlock  uint64
	wasmModule    wasm.Module
	wasmArguments []wasm.Argument
	entrypoint    string
	blockIndex    *index.BlockIndex
	tracer        ttrace.Tracer

	instanceCacheEnabled bool
	cachedInstance       wasm.Instance

	// Results
	logs           []string
	logsTruncated  bool
	executionStack []string
}

func NewBaseExecutor(ctx context.Context, moduleName string, initialBlock uint64, wasmModule wasm.Module, cacheEnabled bool, wasmArguments []wasm.Argument, blockIndex *index.BlockIndex, entrypoint string, tracer ttrace.Tracer) *BaseExecutor {
	return &BaseExecutor{
		ctx:                  ctx,
		initialBlock:         initialBlock,
		blockIndex:           blockIndex,
		moduleName:           moduleName,
		wasmModule:           wasmModule,
		instanceCacheEnabled: cacheEnabled,
		wasmArguments:        wasmArguments,
		entrypoint:           entrypoint,
		tracer:               tracer,
	}
}

// var Timer time.Duration
var ErrNoInput = errors.New("no input")

func (e *BaseExecutor) wasmCall(outputGetter execout.ExecutionOutputGetter) (call *wasm.Call, err error) {
	e.logs = nil
	e.logsTruncated = false
	e.executionStack = nil

	hasInput := false
	for i, input := range e.wasmArguments {
		switch v := input.(type) {
		case *wasm.StoreWriterOutput:
		case *wasm.StoreReaderInput:
			hasInput = true
		case *wasm.ParamsInput:
			hasInput = true
		case wasm.ValueArgument:
			if !v.Active(outputGetter.Clock().Number) {
				break // skipping input that is not active at this block
			}

			data, _, err := outputGetter.Get(v.Name())
			if err != nil {
				if errors.Is(err, execout.ErrNotFound) {
					break
				}
				return nil, fmt.Errorf("input data for %q, param %d: %w", v.Name(), i, err)
			}
			hasInput = true
			v.SetValue(data)
		default:
			panic("unknown wasm argument type")
		}
	}
	if !hasInput {
		return nil, ErrNoInput
	}

	// This allows us to skip the execution of the VM if there are no inputs.
	// This assumption should either be configurable by the manifest, or clearly documented:
	//  state builders will not be called if their input streams are 0 bytes length (and there is no
	//  state store in read mode)
	if hasInput {
		clock := outputGetter.Clock()
		var inst wasm.Instance

		stats := reqctx.ReqStats(e.ctx)
		//t0 := time.Now()
		call = wasm.NewCall(clock, e.moduleName, e.entrypoint, stats, e.wasmArguments)
		inst, err = e.wasmModule.ExecuteNewCall(e.ctx, call, e.cachedInstance, e.wasmArguments)
		//Timer += time.Since(t0)
		if panicErr := call.Err(); panicErr != nil {
			errExecutor := &ErrorExecutor{
				message:    panicErr.Error(),
				stackTrace: call.ExecutionStack,
			}
			return nil, fmt.Errorf("block %d: module %q: general wasm execution panicked: %w: %s", clock.Number, e.moduleName, ErrWasmDeterministicExec, errExecutor.Error())
		}
		if err != nil {
			if err := e.ctx.Err(); err != nil {
				return nil, fmt.Errorf("block %d: module %q: general wasm execution failed: %w", clock.Number, e.moduleName, err)
			}
			return nil, fmt.Errorf("block %d: module %q: general wasm execution failed: %w: %s", clock.Number, e.moduleName, ErrWasmDeterministicExec, err)
		}
		if e.instanceCacheEnabled {
			if err := inst.Cleanup(e.ctx); err != nil {
				return nil, fmt.Errorf("block %d: module %q: failed to cleanup module: %w", clock.Number, e.moduleName, err)
			}
			e.cachedInstance = inst
		} else {
			if err := inst.Close(e.ctx); err != nil {
				return nil, fmt.Errorf("block %d: module %q: failed to close module: %w", clock.Number, e.moduleName, err)
			}
		}
		e.logs = call.Logs
		e.logsTruncated = call.ReachedLogsMaxByteCount()
		e.executionStack = call.ExecutionStack
	}
	return
}

func (e *BaseExecutor) BlockIndex() *index.BlockIndex {
	return e.blockIndex
}

func (e *BaseExecutor) RunsOnBlock(blockNum uint64) bool {
	return blockNum >= e.initialBlock
}

func (e *BaseExecutor) Close(ctx context.Context) error {
	if e.cachedInstance != nil {
		return e.cachedInstance.Close(ctx)
	}
	return nil
}

func (e *BaseExecutor) lastExecutionLogs() (logs []string, truncated bool) {
	return e.logs, e.logsTruncated
}
func (e *BaseExecutor) lastExecutionStack() []string {
	return e.executionStack
}
