package xormplus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var _ XormPlus = (*engine)(nil)

type engine struct {
	*xorm.Engine
	*builder
	logger *Logger
}

func NewEngine(driverName string, dataSourceName string)(*engine,error) {
	en, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	eng := &engine{
		Engine: en,
		builder: &builder{},
	}
	eng.logger = DefaultLogger()
	en.SetLogger(eng.logger)
	return eng, nil
}

func (x *engine)CreateSession()*session  {
	return NewSession(x.Engine.NewSession(), x.builder)
}

func (x *engine) Fetch(rowsSlicePtr interface{}) error {
	session := x.CreateSession()
	defer session.Close()
	return session.Fetch(rowsSlicePtr)
}

func (x *engine) FetchWithPage(rowsSlicePtr interface{}) (*Pagination, error) {
	session := x.CreateSession()
	defer session.Close()

	return session.FetchWithPage(rowsSlicePtr)
}