package xormplus

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/various/xormplus/utils"
	"reflect"
	"strings"
)

var _ Builder = (*builder)(nil)
var errTableMustBeSupply = errors.New("table name must be supply")

type Builder interface {
	WithSort(*Sort) Builder
	WithPageable(Pageable) Builder
	WithTemplate(Template) Builder
	Table(interface{}) Builder
	Columns([]Column) Builder
	Group([]string) Builder
	Having([]Having) Builder
	Or([]Condition) Builder
	And([]Condition) Builder
	Where([]Condition) Builder
	BuildSQL() (string, error)
	BuildCountSQL() (string, error)
	BuildPlus(inc xorm.Interface) (XormPlus, error)
	Clone()Builder
	Clear()
}

type builder struct {
	table    string
	columns  []string
	pageable Pageable
	group    string
	having   string
	sort     string
	or       Query
	and      Query
	where    Query
	query    string
}

func NewBuilder() Builder {
	return &builder{}
}

func (b *builder) WithSort(sort *Sort) Builder {
	b.sort = sort.Inst()
	return b
}

func (b *builder) WithPageable(p Pageable) Builder {
	b.pageable = p
	return b
}

func (b *builder) WithTemplate(template Template) Builder {
	panic("implement me")
}

func (b *builder) Table(bean interface{}) Builder {

	switch t := bean.(type) {
	case string:
		b.table = t
	case Table:
		b.table = t.TableName()
	default:
		tp := reflect.TypeOf(bean)
		if tp.Kind() == reflect.Ptr {
			tp = tp.Elem()
		}
		if tp.Kind() == reflect.Struct {
			b.table = utils.TransCamelToUnderline(tp.Name())
		}
	}

	return b
}

func (b *builder) Columns(columns []Column) Builder {
	for _, column := range columns {
		b.columns = append(b.columns, column.Inst())
	}
	return b
}

func (b *builder) Or(conditions []Condition) Builder {
	if len(conditions) == 0 {
		return b
	}
	var or []string
	for _, condition := range conditions {
		or = append(or, condition.Inst())
	}
	b.or = append(b.or, strings.Join(or, " AND "))
	return b
}

func (b *builder) And(conditions []Condition) Builder {
	if len(conditions) == 0 {
		return b
	}
	var and []string
	for _, condition := range conditions {
		and = append(and, condition.Inst())
	}
	b.and = append(b.and, strings.Join(and, " AND "))
	return b
}

func (b *builder) Where(conditions []Condition) Builder {
	if len(conditions) == 0 {
		return b
	}
	var where []string
	for _, condition := range conditions {
		where = append(where, condition.Inst())
	}
	b.where = append(b.where, strings.Join(where, " AND "))
	return b
}

func (b *builder) Group(group []string) Builder {
	var g []string
	for _, s := range group {
		g = append(g, fmt.Sprintf("`%s`", s))
	}
	b.group = strings.Join(g, ", ")
	return b
}

func (b *builder) Having(havings []Having) Builder {
	var having []string
	for _, hv := range havings {
		having = append(having, hv.Inst())
	}
	b.having = strings.Join(having, " AND ")
	return b
}

func (b *builder) BuildSQL() (string, error) {
	if b.table == "" {
		return "", errTableMustBeSupply
	}
	if _, err := b.buildQuery(); err != nil {
		return "", err
	}

	sb := &strings.Builder{}
	sb.WriteString("SELECT ")
	if len(b.columns)>0{
		sb.WriteString(strings.Join(b.columns, ""))
	}else {
		sb.WriteString("*")
	}
	sb.WriteString(" FROM ")
	sb.WriteString(b.table)

	if b.query != ""{
		sb.WriteString(" WHERE ")
		sb.WriteString(b.query)
	}
	if b.group != "" {
		sb.WriteString(" GROUP BY ")
		sb.WriteString(b.group)
		sb.WriteString(" ")
	}
	if b.having != "" {
		sb.WriteString(" Having ")
		sb.WriteString(b.having)
	}
	if b.sort != ""{
		sb.WriteString(" ORDER BY ")
		sb.WriteString(b.sort)
	}

	if b.pageable != nil{
		sb.WriteString(fmt.Sprintf(" LIMIT %d, %d", b.pageable.Skip(), b.pageable.Limit()) )
	}
	return sb.String(), nil
}

func (b *builder) BuildPlus(inc xorm.Interface) (XormPlus, error) {
	if b.table == "" {
		return nil, errTableMustBeSupply
	}
	if _, err := b.buildQuery(); err != nil {
		return nil, err
	}

	return &xormplus{
		inc: inc,
		builder: b,
		pageable: b.pageable,
	}, nil
}

func (b *builder) Clone()(bc  Builder){
	if b != nil{
		func(b builder) {
			bc = &b
		}(*b)
	}
	return bc
}
func (b *builder) Clear(){
	*b = builder{}
}
func (b *builder) BuildCountSQL()(string, error) {
	if _, err := b.buildQuery(); err != nil {
		return "", err
	}
	if b.query != ""{
		return fmt.Sprintf("SELECT COUNT(1) total FROM %s WHERE %s", b.table, b.query), nil
	}

	return fmt.Sprintf("SELECT COUNT(1) total FROM %s", b.table), nil
}

func (b *builder) buildQuery()(string, error) {
	if b.query != ""{
		return b.query, nil
	}

	var sb strings.Builder
	if len(b.where) > 0 {
		sb.WriteString(b.where.Inst(QueryOpAnd))
	}
	// and ( onw and two ) and (three)
	if len(b.and) > 0 {
		if len(b.where) > 0 {
			sb.WriteString(" AND ")
		}
		sb.WriteString(b.and.Inst(QueryOpAnd))
	}

	// or ( onw and two ) or (three)
	if len(b.or) > 0 {
		if len(b.where) > 0 || len(b.and) > 0 {
			sb.WriteString(" OR ")
		}
		sb.WriteString(b.or.Inst(QueryOpOr))
	}

	b.query = sb.String()

	return b.query, nil
}
