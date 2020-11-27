package xormplus

import (
	"testing"
)

func TestNewSort(t *testing.T) {
	sort := NewSort()
	sort.Column("name")
	sort.Column("id")
	sort.ColumnDesc("tag")
	t.Log(sort.Inst())

	sort.Clear()
	t.Log(sort.Inst())


	sort.FromSlice([]string{"id", "name", "tag DESC"})
	t.Log(sort.Inst())
}
