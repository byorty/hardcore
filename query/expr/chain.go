package expr

import (
	"fmt"
	"github.com/byorty/hardcore/types"
	"strings"
)

type LogicChainImpl struct {
	kind   types.LogicChainKind
	logics []types.Logic
}

func NewLogicChain(kind types.LogicChainKind) types.LogicChain {
	return &LogicChainImpl{
		kind: kind,
	}
}

func (l *LogicChainImpl) WriteSqlPart(writer types.SqlQueryWriter, proto types.Proto, table string, i int) string {
	parts := make([]string, len(l.logics))
	for j, logic := range l.logics {
		var tpl string
		if i == 0 {
			tpl = "%s"
		} else if l.kind == types.AndLogicChainKind {
			tpl = "AND %s"
		} else if l.kind == types.OrLogicChainKind {
			tpl = "OR %s"
		}
		parts[j] = fmt.Sprintf(tpl, logic.WriteSqlPart(writer, proto, table, j))
	}
	return strings.Join(parts, " ")
}

func (l *LogicChainImpl) Add(logic types.Logic) {
	l.logics = append(l.logics, logic)
}
