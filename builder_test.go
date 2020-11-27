package xormplus

import (
	"fmt"
	"github.com/various/xormplus/utils"
	"testing"
)

type CompanyName struct {

}
type User struct {
}
type Xuser User

func (u *User)TableName()string  {
	return "t_user"
}

func TestBuilder_Table(t *testing.T) {
	b := NewBuilder()
	b.Table(&User{})
	b.Table([]string{})
	t.Log(b.BuildSQL())
}
func TestBuilder_Or(t *testing.T) {
	b := NewBuilder()
	b.Table(&User{})

	columns := []Column{
		{
			Name:  "name",
			Alias: "",
		},
		{
			Name:  "id",
			Alias: "pk",
		},
		{
			Name:  "total",
			Func: "COUNT",
			Alias: "total",
		},
	}
	b.Columns(columns)
	or1 := Condition{
		Key:      "name",
		Value:    "select",
		Operator: OperatorEq,
	}
	var o2 uint64 = 1231
	or2 := Condition{
		Key:      "or2",
		Value:    &o2,
		Operator: OperatorEq,
	}
	o3val := "oc3"
	or3 := Condition{
		Key:      "or3",
		Value:    &o3val,
		Operator: OperatorEq,
	}
	var o4val rune =  123
	or4 := Condition{
		Key:      "or4",
		Value:    &o4val,
		Operator: OperatorEq,
	}
	var o5val float64 =  123.0
	or5 := Condition{
		Key:      "or5",
		Value:    &o5val,
		Operator: OperatorEq,
	}
	o6val := []interface{}{1,2,3,34}
	or6 := Condition{
		Key:      "or6",
		Value:    &o6val,
		Operator: OperatorIn,
	}
	o7val := []interface{}{"x","p","x"}
	or7 := Condition{
		Key:      "or7",
		Value:    &o7val,
		Operator: OperatorNotIn,
	}
	o8val := []int{7,8}
	or8 := Condition{
		Key:      "or8",
		Value:    &o8val,
		Operator: OperatorBetween,
	}
	o9val := "likec"
	or9 := Condition{
		Key:      "or9",
		Value:    &o9val,
		Operator: OperatorLike,
	}
	o10val := "likec"
	or10 := Condition{
		Key:      "or10",
		Value:    &o10val,
		Operator: OperatorLike,
	}
	if 10 < 1{
		fmt.Println(or1,or2, or3,or4,or5, or6,or7,or8,or10,or9)
	}

	b.Or([]Condition{or1,or2})
	b.Or([]Condition{or3,or4})
	b.And([]Condition{or8}).And([]Condition{or9})
	b.Where([]Condition{or10,or9})
	b.And([]Condition{or5}).And([]Condition{or6}).Where([]Condition{or9})
	b.Where([]Condition{or7, or8})
	b.Group([]string{"id","name"})
	b.Having([]Having{{
		Key:      "name",
		Value:    10,
		Operator: OperatorGte,
		Func:      "COUNT",
	}, {
		Key:      "id",
		Value:    10,
		Operator: OperatorGte,
		Func:      "SUM",
	}})
	sort := NewSort()
	sort.Column("name")
	sort.ColumnDesc("created_at")
	b.WithSort(sort)
	b.WithPageable(&pageable{
		page:  1,
		size:  39,
	})
	t.Log(b.BuildSQL())
}

func TestBuilder_Build(t *testing.T) {
	s := "xxx \\x alter grant locker drop create"
	t.Log(utils.Filter(s))
}