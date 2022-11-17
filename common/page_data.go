package common

type OrderEnum string

var (
	ASC  OrderEnum = "ASC"
	DESC OrderEnum = "DESC"
)

type PageData[T any] struct {
	PageNo             int       `json:"pageNo"`
	PageSize           int       `json:"pageSize"`
	Order              OrderEnum `json:"order,omitempty"`
	OrderBy            string    `json:"orderBy,omitempty"`
	List               []T       `json:"list"`
	Total              int       `json:"total"`
	TotalPages         int       `json:"totalPages"`
	NeedlessData       bool      `json:"-"`
	NeedlessTotalCount bool      `json:"-"`
}

func (d *PageData[T]) GetFirst() int {
	return (d.PageNo - 1) * d.PageSize
}

func (d *PageData[T]) GetTotalPages() int {
	if d.Total < 0 {
		return -1
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
