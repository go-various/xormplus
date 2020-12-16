package xormplus

import "github.com/go-xorm/xorm"

//xormplus
type XormPlus interface {
	xorm.Interface
	GetById(id interface{}, beanPtr interface{}) (bool, error)
	UpdateById(id interface{}, beanPtr interface{}) (int64, error)
	DeleteById(id interface{}, beanPtr interface{}) (int64, error)

	Fetch(rowsSlicePtr interface{}) error

	FetchWithPage(rowsSlicePtr interface{}) (*Pagination, error)
	NewSession() *session
}
