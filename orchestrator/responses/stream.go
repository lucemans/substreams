package responses

import (
	"github.com/streamingfast/substreams"
	pbsubstreamsrpc "github.com/streamingfast/substreams/pb/sf/substreams/rpc/v2"
)

type Stream struct {
	respFunc substreams.ResponseFunc
}

func New(respFunc substreams.ResponseFunc) *Stream {
	return &Stream{
		respFunc: respFunc,
	}
}

func (s *Stream) InitialProgressMessages(in []*pbsubstreamsrpc.ModuleProgress) {
	s.respFunc(substreams.NewModulesProgressResponse(in))
}

/*
outputstream.Walker
orchestrator/execout/stream.go Stream
orchestrator/execout/walker.go Walker
orchestrator/linear/reader.go Reader
orchestrator/execout/linearreader.go LinearReader
orchestrator/execout/walker.go execout.Walker


*/
