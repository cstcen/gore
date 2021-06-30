package model

type PageResult struct {
	BaseResult
	Data []PageData `json:"data"`
}
