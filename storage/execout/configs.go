package execout

import (
	"fmt"

	"github.com/streamingfast/dstore"
	"github.com/streamingfast/substreams/block"
	"github.com/streamingfast/substreams/manifest"
	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
	"go.uber.org/zap"
)

type Configs struct {
	ConfigMap              map[string]*Config
	execOutputSaveInterval uint64
	logger                 *zap.Logger
}

func NewConfigs(baseObjectStore dstore.Store, allRequestedModules []*pbsubstreams.Module, moduleHashes *manifest.ModuleHashes, execOutputSaveInterval uint64, logger *zap.Logger) (*Configs, error) {
	out := make(map[string]*Config)
	for _, mod := range allRequestedModules {
		conf, err := NewConfig(
			mod.Name,
			mod.InitialBlock,
			mod.ModuleKind(),
			moduleHashes.Get(mod.Name),
			baseObjectStore,
			logger,
		)
		if err != nil {
			return nil, fmt.Errorf("new exec output config for %q: %w", mod.Name, err)
		}
		out[mod.Name] = conf
	}

	return &Configs{
		execOutputSaveInterval: execOutputSaveInterval,
		ConfigMap:              out,
		logger:                 logger,
	}, nil
}

func (c *Configs) NewFile(moduleName string, targetRange *block.Range) *File {
	return c.ConfigMap[moduleName].NewFile(targetRange)
}

func (c *Configs) NewFileWalker(moduleName string, segmenter *block.Segmenter) *FileWalker {
	return NewFileWalker(c.ConfigMap[moduleName], segmenter, c.logger)
}
