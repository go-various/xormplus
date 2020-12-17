package xormplus

import (
	"github.com/hashicorp/go-hclog"
	"testing"
)

type Banner struct {
	Id         *int    `json:"id" xorm:"id autoincr not null pk comment('pk') INT"`
	Title      *string `json:"title" xorm:"title not null comment('标题') VARCHAR(64)"`
}

func (o *Banner) TableName() string {
	return "t_banner"
}
func (o *Banner) ID() interface{} {
	return o.Id
}

func (o *Banner) NewEntity(dao XormPlus) Entity {
	return NewEntity(dao, o)
}

func  TestEntity_Detail(t *testing.T) {
	plus, err := NewEngine(&Config{
		ShowSql:        true,
		MaxIdle:        0,
		MaxConn:        0,
		Master:         "root:123456@tcp(127.0.0.1)/coin?charset=utf8mb4&parseTime=true&loc=Local",
		Slaves:         nil,
		UseMasterSlave: false,
	}, hclog.Default())

	if nil != err{
		t.Fatal(err)
	}
	banner := &Banner{}

	session := plus.NewSession()
	session.Begin()
	defer session.Close()

	entity := banner.NewEntity(session)
	has, err := entity.Exists()
	if nil != err{
		t.Fatal(err)
	}

	t.Log(has, entity.Entity())
}