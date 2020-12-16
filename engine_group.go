package xormplus

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/hashicorp/go-hclog"
	"xorm.io/core"
)

var _ XormPlus = (*engineGroup)(nil)

type engineGroup struct {
	*xorm.EngineGroup
	*builder
	logger *Logger
}

func NewEngineGroup(c *Config, logger hclog.Logger) (*engineGroup, error) {

	if nil == c || "" == c.Master {
		return nil, errors.New("config or config.Url can not be null")
	}
	conns := make([]string, len(c.Slaves)+1)
	conns[0] = c.Master
	for i, v := range c.Slaves {
		conns[i+1] = v
		if "" == v {
			return nil, errors.New("config or config.Url can not be null")
		}
	}

	group, err := xorm.NewEngineGroup("xorm", conns)

	if nil != err || nil == group {
		return nil, err
	}
	log := NewLogger(logger)

	group.SetLogger(log)
	group.SetMapper(core.GonicMapper{})
	group.ShowSQL(c.ShowSql)
	group.SetMaxIdleConns(c.MaxIdle)
	group.SetMaxOpenConns(c.MaxConn)

	eng := &engineGroup{
		EngineGroup: group,
		builder:     &builder{},
		logger:      log,
	}

	return eng, nil
}
func (x *engineGroup) NewSession() *session {
	return NewSession(x.Engine.NewSession(), x.builder)
}
func (x *engineGroup) Fetch(rowsSlicePtr interface{}) error {
	session := x.NewSession()
	defer session.Close()
	return session.Fetch(rowsSlicePtr)
}

func (x *engineGroup) FetchWithPage(rowsSlicePtr interface{}) (*Pagination, error) {
	session := x.NewSession()
	defer session.Close()

	return session.FetchWithPage(rowsSlicePtr)
}

func (x *engineGroup) GetById(id interface{}, beanPtr interface{}) (bool, error) {
	session := x.NewSession()
	defer session.Close()
	return session.GetById(id, beanPtr)
}

func (x *engineGroup) UpdateById(id interface{}, beanPtr interface{}) (int64, error) {
	session := x.NewSession()
	defer session.Close()
	return session.UpdateById(id, beanPtr)
}

func (x *engineGroup) DeleteById(id interface{}, beanPtr interface{}) (int64, error) {
	session := x.NewSession()
	defer session.Close()
	return session.DeleteById(id, beanPtr)
}
