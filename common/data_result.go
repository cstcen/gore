package common

type DataResult[T any] struct {
	BaseResult
	Data T `json:"data,omitempty"`
}

func NewDataResult[T any](data T) *DataResult[T] {
	return &DataResult[T]{BaseResult: BaseResult{ErrSuccess}, Data: data}
}
