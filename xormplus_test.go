package xormplus

import (
	"github.com/go-xorm/xorm"
	"testing"
)

type Task struct {
	TaskId int
	TaskName string
}

func (a Task)TableName()string  {
	return "t_task_log"
}

func TestXormplus_NewSession(t *testing.T) {
	engine, err := xorm.NewEngine("mysql",
		"root:123456@tcp(127.0.0.1)/test?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		t.Fatal(err)
		return
	}
	var out []Task

	engine.ShowSQL(true)
	b := NewBuilder()
	b.Where([]Condition{{
		Key:      "task_name",
		Value:    "test-job",
		Operator: OperatorEq,
	}})
	b.And([]Condition{
		{
			Key:      "time_created",
			Value:    "2020-03-06 16:22:03",
			Operator: OperatorGte,
		},
		{
			Key:      "time_created",
			Value:    "2020-03-06 16:22:03",
			Operator: OperatorLt,
		},
	})
	b.Or([]Condition{{
		Key:      "task_name",
		Value:    "test-job",
		Operator: OperatorEq,
	}})
	b.Table(&Task{})

	b.Columns([]Column{{
		Name:  "task_name",
		Func:  "",
		Alias: "",
	}})
	b.WithPageable(NewPageable(0,5))
	b.Group([]string{"task_id"})
	t.Log(b.BuildSQL())
	sort := NewSort().Column("task_id").ColumnDesc("time_created")
	b.WithSort(sort)
	b.Having([]Having{{
		Key:      "task_id",
		Value:    1,
		Operator: OperatorGt,
		Func:     "COUNT",
	}})
	plus, err := b.BuildPlus(engine.NewSession())
	if err != nil {
		t.Fatal(err)
	}


	if pg, err := plus.FetchWithPage(&out); err != nil {
		t.Fatal(err)
		return
	}else {
		t.Log(pg)
	}
	t.Log(out)

}
