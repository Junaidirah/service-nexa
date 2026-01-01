package helper

type Meta struct {
	Pagination *Pagination `json:"pagination,omitempty"`
}
type Pagination struct {
	TotalCount int `json:"total_count"`
	TotalPages int `json:"total_pages"`
	PerPage    int `json:"per_page"`
	Page       int `json:"page"`
}
type ApiResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Error   error  `json:"error"`
	Meta    *Meta  `json:"meta,omitempty"`
}

func SuccessWithMeta[T any](data T, meta Meta, message string) ApiResponse[T] {
	return ApiResponse[T]{
		Success: true,
		Data:    data,
		Meta:    &meta,
		Message: message,
	}
}

func Error(err error, message string) ApiResponse[any] {
	return ApiResponse[any]{
		Success: false,
		Error:   err,
		Message: message,
	}
}
