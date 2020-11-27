package xormplus

import (
	"fmt"
	"reflect"
)

type Operator string

const (
	OperatorNe         Operator = "<>"
	OperatorEq         Operator = "="
	OperatorLt         Operator = "<"
	OperatorGt         Operator = ">"
	OperatorLte        Operator = ">="
	OperatorGte        Operator = ">="
	OperatorIn         Operator = "IN"
	OperatorNotIn      Operator = "NOT IN"
	OperatorLike       Operator = "LIKE"
	OperatorNotLike    Operator = "NOT LIKE"
	OperatorBetween    Operator = "BETWEEN"
	OperatorNotBetween Operator = "NOT BETWEEN"
)

func scalarValue(in interface{}) interface{} {
	switch reflect.TypeOf(in).Kind() {
	case reflect.Ptr:
		value := reflect.ValueOf(in).Elem()
		return scalarValue(value.Interface())
	case reflect.String:
		return fmt.Sprintf("'%v'", in)
	default:
		return fmt.Sprintf("%v", in)
	}
}
func sliceValue(in interface{})[]interface{}  {
	var out []interface{}
	s := reflect.ValueOf(in).Elem()
	for i := 0; i < s.Len(); i++ {
		out = append(out, scalarValue(s.Index(i).Interface()))
	}
	return out
}