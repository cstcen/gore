package common

type OrderEnum string

var (
	ASC  OrderEnum = "ASC"
	DESC OrderEnum = "DESC"
)

type PageData[T any] struct {
	PageNo             int       `json:"pageNo,omitempty"`
	PageSize           int       `json:"pageSize,omitempty"`
	Order              OrderEnum `json:"order,omitempty"`
	OrderBy            string    `json:"orderBy,omitempty"`
	List               []T       `json:"list"`
	Total              int       `json:"total,omitempty"`
	TotalPages         int       `json:"totalPages,omitempty"`
	NeedlessData       bool      `json:"-"`
	NeedlessTotalCount bool      `json:"-"`
}

func NewPageData[T any](list []T) *PageData[T] {
	return &PageData[T]{List: list}
}

func (d *PageData[T]) WithTotal(total int) *PageData[T] {
	d.Total = total
	d.TotalPages = d.GetTotalPages()
	return d
}

func (d *PageData[T]) WithPageNo(pageNo int) *PageData[T] {
	d.PageNo = pageNo
	return d
}

func (d *PageData[T]) WithPageSize(pageSize int) *PageData[T] {
	d.PageSize = pageSize
	d.TotalPages = d.GetTotalPages()
	return d
}

func (d *PageData[T]) GetFirst() int {
	return (d.PageNo - 1) * d.PageSize
}

func (d *PageData[T]) GetTotalPages() int {
	if d.Total <= 0 || d.PageSize == 0 {
		return 0
	}

	count := d.Total / d.PageSize

	if d.Total%d.PageSize > 0 {
		count++
	}

	return count
}

func (d *PageData[T]) HasNext() bool {
	return d.PageNo+1 <= d.GetTotalPages()
}

func (d *PageData[T]) NextPage() int {
	if d.HasNext() {
		return d.PageNo + 1
	}
	return d.PageNo
}

func (d *PageData[T]) EndIndex() int {
	return d.PageNo * d.PageSize
}

func (d *PageData[T]) BeginIndex() int {
	return d.GetFirst()
}

func (d *PageData[T]) HasPrev() bool {
	return d.PageNo-1 >= 1
}

func (d *PageData[T]) PrevPage() int {
	if d.HasPrev() {
		return d.PageNo - 1
	}
	return d.PageNo
}
