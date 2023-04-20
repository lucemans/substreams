package ranges

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	pbsubstreamsrpc "github.com/streamingfast/substreams/pb/sf/substreams/rpc/v2"
	"github.com/streamingfast/substreams/tui2/common"
)

// Bar is a single progress bar that handles ranges with holes

type Bar struct {
	common.Common
	name           string
	targetEndBlock uint64

	ranges ranges
}

func NewBar(c common.Common, name string, targetEndBlock uint64) *Bar {
	out := &Bar{Common: c, name: name, targetEndBlock: targetEndBlock}

	return out
}

func (b *Bar) Init() tea.Cmd { return nil }

func (b *Bar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case *pbsubstreamsrpc.ModuleProgress_ProcessedRanges:
		for _, v := range msg.ProcessedRanges {
			b.ranges = mergeRangeLists(b.ranges, &blockRange{
				Start: v.StartBlock,
				End:   v.EndBlock,
			})
		}
	}
	return b, nil
}

func (b *Bar) View() string {
	width := uint64(b.Width)
	if width > 1000 {
		return "[borked]"
	}

	in := b.ranges
	lo := in.Lo()
	hi := b.targetEndBlock
	binSize := (hi - lo) / width
	var out []string
	for i := uint64(0); i < width; i++ {
		loCheck := binSize*i + lo
		hiCheck := binSize*(i+1) + lo

		if in.Covered(loCheck, hiCheck) {
			out = append(out, "▓")
		} else if in.PartiallyCovered(loCheck, hiCheck) {
			out = append(out, "▒")
		} else {
			out = append(out, "░")
		}
	}
	return "[" + strings.Join(out, "") + "]"
}

func (b *Bar) RangeView() string {
	width := uint64(b.Width)
	if width > 1000 {
		return "[borked]"
	}

	in := b.ranges
	lo := in.Lo()
	hi := b.targetEndBlock
	var out []string
	out = append(out, fmt.Sprintf("%v - %v", lo, hi))

	return "[" + strings.Join(out, "") + "]"
}
