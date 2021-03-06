package xormplus

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/hashicorp/go-hclog"
	"xorm.io/core"
)

var _ XormPlus = (*engine)(nil)

type engine struct {
	*xorm.Engine
	*builder
	logger *Logger
}

func NewEngine(c *Config, logger hclog.Logger) (*engine, error) {
	if nil == c || "" == c.Master {
		return nil, errors.New("config or config.Url can not be null")
	}
	en, err := xorm.NewEngine("mysql", c.Master)
	if err != nil {
		return nil, err
	}
	eng := &engine{
		Engine:  en,
		builder: &builder{},
	}

	eng.logger = NewLogger(logger)
	en.SetLogger(eng.logger)
	eng.ShowSQL(c.ShowSql)
	eng.SetMapper(core.GonicMapper{})
	eng.SetMaxIdleConns(c.MaxIdle)
	eng.SetMaxOpenConns(c.MaxConn)

	return eng, nil
}

func (e *engine) NewSession() *Session {
	return NewSession(e.Engine.NewSession(), e.builder)
}

func (e *engine) Fetch(rowsSlicePtr interface{}) error {
	session := e.NewSession()
	defer session.Close()
	return session.Fetch(rowsSlicePtr)
}

func (e *engine) FetchWithPage(rowsSlicePtr interface{}) (*Pagination, error) {
	session := e.NewSession()
	defer session.Close()

	return session.FetchWithPage(rowsSlicePtr)
}

func (e *engine) GetById(id interface{}, beanPtr interface{}) (bool, error) {
	session := e.NewSession()
	defer session.Close()
	return session.GetById(id, beanPtr)
}

func (e *engine) UpdateById(id interface{}, beanPtr interface{}) (int64, error) {
	session := e.NewSession()
	defer session.Close()
	return session.UpdateById(id, beanPtr)
}

func (e *engine) DeleteById(id interface{}, beanPtr interface{}) (int64, error) {
	session := e.NewSession()
	defer session.Close()
	return session.DeleteById(id, beanPtr)
}
