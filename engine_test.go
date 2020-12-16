package xormplus

import (
	"github.com/hashicorp/go-hclog"
	"testing"
)

type Task struct {
	TaskId   int
	TaskName string
}

func (a Task) TableName() string {
	return "t_task_log"
}

func TestXormplus_NewSession(t *testing.T) {
	b, err := NewEngine(&Config{}, hclog.Default())
	if err != nil {
		t.Fatal(err)
		return
	}
	b.NewSession().Close()
	b.WithCondition([]Condition{{
		Key:      "task_name",
		Value:    "test-job",
		Operator: OperatorEq,
	}})
	b.WithAnd([]Condition{
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
	b.WithOr([]Condition{{
		Key:      "task_name",
		Value:    "test-job",
		Operator: OperatorEq,
	}})
	b.WithTable(&Task{})

	b.WithColumns([]Column{{
		Name:  "task_name",
		Func:  "",
		Alias: "",
	}})
	b.WithPageable(NewPageable(0, 5))
	b.WithGroup([]string{"task_id"})

	sort := NewSort().Column("task_id").ColumnDesc("time_created")
	b.WithSort(sort)
	b.WithHaving([]Having{{
		Key:      "task_id",
		Value:    1,
		Operator: OperatorGt,
		Func:     "COUNT",
	}})
	t.Log(b.BuildSQL())
	var tasks []Task
	if p, err := b.FetchWithPage(&tasks); err != nil {
		t.Fatal(err)
		return
	} else {
		t.Log(p)
	}
	session := b.NewSession()
	defer session.Close()
	session.Fetch(&tasks)
}
