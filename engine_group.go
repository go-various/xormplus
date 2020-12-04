package xormplus

import (
	"github.com/go-xorm/xorm"
)

var _ XormPlus = (*engineGroup)(nil)

type engineGroup struct {
	*xorm.EngineGroup
	*builder
	logger *Logger
}

func NewEngineGroup(args1 interface{}, args2 interface{}, policies ...xorm.GroupPolicy)(*engineGroup,error) {
	en, err := xorm.NewEngineGroup(args1, args2, policies...)
	if err != nil {
		return nil, err
	}
	eng := &engineGroup{
		EngineGroup: en,
		builder: &builder{},
	}
	eng.logger = DefaultLogger()
	en.SetLogger(eng.logger)
	return eng, nil
}
func (x *engineGroup)CreateSession()*session  {
	return NewSession(x.Engine.NewSession(), x.builder)
}

func (x *engineGroup) Fetch(rowsSlicePtr interface{}) error {
	session := x.CreateSession()
	defer session.Close()
	return session.Fetch(rowsSlicePtr)
}

func (x *engineGroup) FetchWithPage(rowsSlicePtr interface{}) (*Pagination, error) {
	session := x.CreateSession()
	defer session.Close()

	return session.FetchWithPage(rowsSlicePtr)
}