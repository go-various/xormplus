//查询条件
package xormplus

import "fmt"

type Condition struct {
	Key      string
	Value    interface{}
	Operator Operator
}

func (c Condition) Inst() string {
	switch c.Operator {
	case OperatorIn, OperatorNotIn:
		values := sliceValue(c.Value)
		if len(values) == 0{
			return ""
		}
		in := ""
		for _, val := range values {
			in = fmt.Sprintf("%s,%v", in, val)
		}
		return fmt.Sprintf("`%s` %s (%v)", c.Key, c.Operator, in[1:])
	case OperatorBetween, OperatorNotBetween:
		values := sliceValue(c.Value)
		if len(values) != 2{
			return ""
		}
		return fmt.Sprintf("`%s` %s %v AND %v", c.Key, c.Operator, values[0], values[1])
	default:
		return fmt.Sprintf("`%s` %s %v", c.Key, c.Operator, scalarValue(c.Value))
	}
}
