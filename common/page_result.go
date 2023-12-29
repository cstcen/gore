package common

type PageResult[T any] struct {
	BaseResult
	Data PageData[T] `json:"data,omitempty"`
}

func NewPageResult[T any](data PageData[T]) *PageResult[T] {
	return &PageResult[T]{BaseResult: BaseResult{ErrSuccess}, Data: data}
}
