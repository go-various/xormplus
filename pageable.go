//分页
package xormplus

type Pageable interface {
	Skip() int
	Limit() int
	Page() int
	Size() int
}

type pageable struct {
	page  int
	size  int
}

func (p *pageable) Skip() int {
	return p.page * p.size
}
func (p *pageable) Limit() int {
	return p.size
}

func (p *pageable) Size() int {
	return p.size
}
func (p *pageable) Page() int {
	return p.page
}

func NewPageable(page, size int) Pageable {
	return &pageable{page: page, size: size}
}

func DefaultPageable() Pageable {
	return &pageable{page: 0, size: 50}
}