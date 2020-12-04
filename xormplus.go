package xormplus

import "github.com/go-xorm/xorm"

//xormplus
type XormPlus interface {
	xorm.EngineInterface

	Fetch(rowsSlicePtr interface{}) error

	FetchWithPage(rowsSlicePtr interface{}) (*Pagination, error)
}