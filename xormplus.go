package xormplus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strconv"
)

//xorm增强查询工具
//转换到xorm的查询方法为:  engine.SQL(x.sql).XX(beanPtr)
type XormPlus interface {
	//列表查询
	Fetch(rowsSlicePtr interface{}) error
	//列表查询分页
	FetchWithPage(rowsSlicePtr interface{}) (*Pagination, error)
}

type xormplus struct {
	inc xorm.Interface
	builder Builder
	pageable Pageable
}

func (x *xormplus) Fetch(rowsSlicePtr interface{}) error {
	sql, err := x.builder.BuildSQL()
	if err != nil {
		return err
	}
	return x.inc.SQL(sql).Find(rowsSlicePtr)
}

func (x *xormplus) FetchWithPage(rowsSlicePtr interface{}) (*Pagination, error) {

	query, err := x.builder.BuildCountSQL()
	if err != nil {
		return nil, err
	}

	total, err := x.inc.QueryString(query)
	if err != nil {
		return nil, err
	}

	sql, err := x.builder.BuildSQL()
	if err != nil {
		return nil, err
	}
	err = x.inc.SQL(sql).Find(rowsSlicePtr)
	if err != nil {
		return nil, err
	}

	tt, _ := strconv.Atoi(total[0]["total"])
	pg := NewPagination(tt, x.pageable)
	return pg, nil
}
