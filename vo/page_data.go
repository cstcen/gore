package vo

type OrderEnum string

var (
	ASC  OrderEnum = "ASC"
	DESC OrderEnum = "DESC"
)

type PageData struct {
	PageNo             int           `json:"pageNo,omitempty"`
	PageSize           int           `json:"pageSize,omitempty"`
	Order              OrderEnum     `json:"order,omitempty"`
	OrderBy            string        `json:"orderBy,omitempty"`
	NeedlessData       bool          `json:"-"`
	NeedlessTotalCount bool          `json:"-"`
	List               []interface{} `json:"list,omitempty"`
	Total              int           `json:"total,omitempty"`
}

func (d *PageData) GetFirst() int {
	return (d.PageNo - 1) * d.PageSize
}

func (d *PageData) GetTotalPages() int {
	if d.Total < 0 {
		return -1
	}

	count := d.Total / d.PageSize

	if d.Total%d.PageSize > 0 {
		count++
	}

	return count
}

func (d *PageData) HasNext() bool {
	return d.PageNo+1 <= d.GetTotalPages()
}

func (d *PageData) NextPage() int {
	if d.HasNext() {
		return d.PageNo + 1
	}
	return d.PageNo
}

func (d *PageData) EndIndex() int {
	return d.PageNo * d.PageSize
}

func (d *PageData) BeginIndex() int {
	return d.GetFirst()
}

func (d *PageData) HasPrev() bool {
	return d.PageNo-1 >= 1
}

func (d *PageData) PrevPage() int {
	if d.HasPrev() {
		return d.PageNo - 1
	}
	return d.PageNo
}
