//数据表列
package xormplus

import "fmt"

type Column struct {
	Name  string
	Func  string
	Alias string
}

func (c Column)Inst()string  {
	name := fmt.Sprintf("`%s`", c.Name)
	if c.Func != "" {
		name = fmt.Sprintf("%s(`%s`)", c.Func, c.Name)
	}
	if c.Alias != "" {
		return fmt.Sprintf("%s AS %s", name, c.Alias)
	}
	return name
}