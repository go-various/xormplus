package xormplus

import (
	"errors"
	"github.com/go-xorm/xorm"
	"reflect"
	"strconv"
)

var _ XormPlus = (*Session)(nil)

type Session struct {
	*xorm.Session
	Builder
}

func (session *Session) NewSession() *Session {
	return &Session{Session: session.Session, Builder: session.Builder}
}

func NewSession(s *xorm.Session, b Builder) *Session {
	return &Session{Session: s, Builder: b}
}
func (session *Session) Begin() error {
	return session.Session.Begin()
}

func (session *Session) Commit() error {
	return session.Session.Commit()
}

func (session *Session) Rollback() error {
	return session.Session.Rollback()
}

func (session *Session) Close() {
	session.Session.Close()
}

func (session *Session) Fetch(rowsSlicePtr interface{}) error {
	sql, err := session.Builder.BuildSQL()
	if err != nil {
		return err
	}
	return session.Session.SQL(sql).Find(rowsSlicePtr)
}

func (session *Session) FetchWithPage(rowsSlicePtr interface{}) (*Pagination, error) {
	if session.Pageable() == nil {
		return nil, errors.New("pageable not supplied")
	}
	query, err := session.Builder.BuildCountSQL()
	if err != nil {
		return nil, err
	}

	total, err := session.Session.QueryString(query)
	if err != nil {
		return nil, err
	}

	sql, err := session.Builder.BuildSQL()
	if err != nil {
		return nil, err
	}
	err = session.Session.SQL(sql).Find(rowsSlicePtr)
	if err != nil {
		return nil, err
	}

	tt, _ := strconv.Atoi(total[0]["total"])
	pg := NewPagination(tt, session.Pageable())
	return pg, nil
}

func (session *Session) GetById(id interface{}, beanPtr interface{}) (bool, error) {
	if reflect.ValueOf(id).IsZero() {
		return false, errors.New("id cannot be nil")
	}
	return session.Session.ID(id).Get(beanPtr)
}

func (session *Session) UpdateById(id interface{}, beanPtr interface{}) (int64, error) {
	if reflect.ValueOf(id).IsZero() {
		return 0, errors.New("id cannot be nil")
	}
	return session.Session.ID(id).Update(beanPtr)
}

func (session *Session) DeleteById(id interface{}, beanPtr interface{}) (int64, error) {
	if reflect.ValueOf(id).IsZero() {
		return 0, errors.New("id cannot be nil")
	}
	return session.Session.ID(id).Delete(beanPtr)
}
