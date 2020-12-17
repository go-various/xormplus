package xormplus

type Table interface {
	TableName()string
	ID() interface{}
}