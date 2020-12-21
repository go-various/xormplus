package xormplus

import (
	"errors"
	"reflect"
)
var errBeanMustBePointer = errors.New("entity bean must be a type of pointer")
var errBeanMustBeStruct = errors.New("entity bean must be a type of ptr to struct")

var errBeanMustBeNotNil = errors.New("entity bean must be not nil")

var _ Entity = (*entity)(nil)


type Entity interface {
	ID() interface{}
	Detail() (bool, error)
	Exists() (bool, error)
	Create() (int64, error)
	Update() (int64, error)
	Remove() (int64, error)
	Object() interface{}
}

type entity struct {
	dao     XormPlus
	beanPtr interface{}
}

func NewEntity(dao XormPlus, beanPtr interface{}) Entity {
	return &entity{dao: dao, beanPtr: beanPtr}
}

func (e *entity) Object() interface{} {
	return e.beanPtr
}

func (e *entity) ID() interface{} {
	return e.beanPtr.(Table).ID()
}

func (e *entity) Exists() (bool, error) {
	if err := checkBean(e.beanPtr); err != nil {
		return false, err
	}
	return e.dao.Exist(e.beanPtr)
}

func (e *entity) Detail() (has bool, err error) {
	if err := checkBean(e.beanPtr); err != nil {
		return false, err
	}
	return e.dao.ID(e.beanPtr.(Table).ID()).Get(e.beanPtr)
}

func (e *entity) Create() (int64, error) {
	if err := checkBean(e.beanPtr); err != nil {
		return 0, err
	}
	return e.dao.InsertOne(e.beanPtr)
}

func (e *entity) Update() (n int64,err error) {
	if err := checkBean(e.beanPtr); err != nil {
		return 0, err
	}
	return e.dao.ID(e.beanPtr.(Table).ID()).Update(e.beanPtr)
}

func (e *entity) Remove() (rows int64, err error) {
	if err := checkBean(e.beanPtr); err != nil {
		return 0, err
	}

	return e.dao.ID(e.beanPtr.(Table).ID()).Delete( e.beanPtr)
}

func checkBean(beanPtr interface{})error  {
	if beanPtr == nil{
		return errBeanMustBeNotNil
	}
	bean := reflect.TypeOf(beanPtr)
	if bean.Kind() != reflect.Ptr{
		return errBeanMustBePointer
	}
	if bean.Elem().Kind() != reflect.Struct{
		return errBeanMustBeStruct
	}
	switch beanPtr.(type) {
	case Table:
		return nil
	default:
		return errors.New("interface conversion: entity bean " +bean.String() + " is not xormplus.Table: missing method ID")
	}
}