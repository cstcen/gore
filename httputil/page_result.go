package httputil

type PageResult struct {
	BaseResult
	Data []PageData `json:"data"`
}
