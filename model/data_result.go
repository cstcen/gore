package model

type DataResult struct {
	BaseResult
	Data interface{} `json:",omitempty"`
}
