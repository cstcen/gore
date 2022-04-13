package vo

type PageResult[T any] struct {
	BaseResult
	Data []PageData[T] `json:"data,omitempty"`
}
