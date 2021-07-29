package model

type DataResult struct {
	BaseResult
	Data interface{} `json:"data,omitempty"`
}
