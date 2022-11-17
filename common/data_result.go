package common

type DataResult[T any] struct {
	BaseResult
	Data T `json:"data,omitempty"`
}
