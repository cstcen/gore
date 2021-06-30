package httputil

type DataResult struct {
	BaseResult
	Data interface{} `json:"data"`
}
