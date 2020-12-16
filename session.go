package xormplus

import (
	"errors"
	"github.com/go-xorm/xorm"
	"reflect"
	"strconv"
)

type session struct {
	*xorm.Session
	Builder
}

func NewSession(s *xorm.Session, b Builder) *session {
	return &session{Session: s, Builder: b}
}
func (x *session) Begin() error {
	return x.Session.Begin()
}
func (x *session) Commit() error {
	return x.Session.Commit()
}

func (x *session) Rollback() error {
	return x.Session.Rollback()
}

func (x *session) Close() {
	x.Session.Close()
}

func (x *session) Fetch(rowsSlicePtr interface{}) error {
	sql, err := x.Builder.BuildSQL()
	if err != nil {
		return err
	}
	return x.Session.SQL(sql).Find(rowsSlicePtr)
}

func (x *session) FetchWithPage(rowsSlicePtr interface{}) (*Pagination, error) {
	if x.Pageable() == nil {
		return nil, errors.New("pageable not supplied")
	}
	query, err := x.Builder.BuildCountSQL()
	if err != nil {
		return nil, err
	}

	total, err := x.Session.QueryString(query)
	if err != nil {
		return nil, err
	}

	sql, err := x.BuildSQL()
	if err != nil {
		return nil, err
	}
	err = x.Session.SQL(sql).Find(rowsSlicePtr)
	if err != nil {
		return nil, err
	}

	tt, _ := strconv.Atoi(total[0]["total"])
	pg := NewPagination(tt, x.Pageable())
	return pg, nil
}

func (x *session) GetById(id interface{}, beanPtr interface{}) (bool, error) {
	if reflect.ValueOf(id).IsZero() {
		return false, errors.New("id cannot be nil")
	}
	return x.Session.ID(id).Get(beanPtr)
}

func (x *session) UpdateById(id interface{}, beanPtr interface{}) (int64, error) {
	if reflect.ValueOf(id).IsZero() {
		return 0, errors.New("id cannot be nil")
	}
	return x.Session.ID(id).Update(beanPtr)
}

func (x *session) DeleteById(id interface{}, beanPtr interface{}) (int64, error) {
	if reflect.ValueOf(id).IsZero() {
		return 0, errors.New("id cannot be nil")
	}
	return x.Session.ID(id).Delete(beanPtr)
}
