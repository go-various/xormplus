//查询条件构造器

package xormplus

import "strings"

type QueryOp string

const (
	QueryOpOr QueryOp = "OR"
	QueryOpAnd QueryOp = "AND"
)

type Query []string

func (o Query)Inst(op QueryOp) string {
	var sb strings.Builder
	for i := 0; i < len(o); i++ {
		sb.WriteString("(")
		sb.WriteString(o[i])
		sb.WriteString(")")
		if i < len(o)-1{
			sb.WriteString(string(op))
		}
	}
	return sb.String()
}