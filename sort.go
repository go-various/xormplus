package xormplus

import (
	"fmt"
	"strings"
)

type Sort struct {
	sorts []string
}

func NewSort() *Sort {
	return &Sort{sorts: make([]string, 0)}
}

func (s *Sort) Column(col string)*Sort{
	s.sorts =  append(s.sorts, fmt.Sprintf("`%s`", col))
	return s
}

func (s *Sort) ColumnDesc(col string)*Sort{
	s.sorts =  append(s.sorts, fmt.Sprintf("`%s` DESC",col))
	return s
}

func (s *Sort)FromSlice(slice []string)*Sort {
	s.sorts = append(s.sorts, slice...)
	return s
}

func (s *Sort) Inst() string {
	return strings.Join(s.sorts, ", ")
}
func (s *Sort) Clear() *Sort{
	s.sorts = []string{}
	return s
}
