package xormplus
// http返回分页数据结构
type Pagination struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func NewPagination(total int, page Pageable) *Pagination {
	mod := total % page.Size()
	totalPages := total / page.Size()
	if 0 == totalPages {
		totalPages = 1
	}
	if mod > 0 && total > page.Size() {
		totalPages += 1
	}
	return &Pagination{Total: total, TotalPages: totalPages, Page: page.Page(), Size: page.Size()}
}