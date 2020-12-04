package xormplus

import (
	"errors"
	"github.com/go-xorm/xorm"
	"strconv"
)

type session struct {
	*xorm.Session
	Builder
}

func NewSession(s *xorm.Session, b Builder) *session {
	return &session{Session: s, Builder: b}
}

func (x *session) Fetch(rowsSlicePtr interface{}) error {
	sql, err := x.BuildSQL()
	if err != nil {
		return err
	}
	return x.SQL(sql).Find(rowsSlicePtr)
}

func (x *session) FetchWithPage(rowsSlicePtr interface{}) (*Pagination, error) {
	if x.Pageable() == nil{
		return nil, errors.New("pageable not supplied")
	}
	query, err := x.BuildCountSQL()
	if err != nil {
		return nil, err
	}

	total, err := x.QueryString(query)
	if err != nil {
		return nil, err
	}

	sql, err := x.BuildSQL()
	if err != nil {
		return nil, err
	}
	err = x.SQL(sql).Find(rowsSlicePtr)
	if err != nil {
		return nil, err
	}

	tt, _ := strconv.Atoi(total[0]["total"])
	pg := NewPagination(tt, x.Pageable())
	return pg, nil
}