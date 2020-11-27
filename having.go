package xormplus

import "fmt"

type Having struct {
	Key      string
	Value    interface{}
	Operator Operator
	Func string
}
func (c Having) Inst() string {
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
		if c.Func != ""{
			return fmt.Sprintf("%s(`%s`) %s (%v)",
				c.Func, c.Key, c.Operator, in[1:])
		}
		return fmt.Sprintf("`%s` %s (%v)", c.Key, c.Operator, in[1:])
	case OperatorBetween, OperatorNotBetween:
		values := sliceValue(c.Value)
		if len(values) != 2{
			return ""
		}
		if c.Func != ""{
			return fmt.Sprintf("%s(`%s`) %s %v AND %v",
				c.Func,c.Key, c.Operator, values[0], values[1])
		}
		return fmt.Sprintf("`%s` %s %v AND %v", c.Key, c.Operator, values[0], values[1])
	default:
		if c.Func != ""{
			return fmt.Sprintf("%s(`%s`) %s %v",
				c.Func,c.Key, c.Operator, scalarValue(c.Value))
		}
		return fmt.Sprintf("`%s` %s %v", c.Key, c.Operator, scalarValue(c.Value))
	}
}