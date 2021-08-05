package vo

type PageResult struct {
	BaseResult
	Data []PageData `json:"data,omitempty"`
}
