package utils

type Pagination struct {
	Limit  int `json:"limit"`
	LastID int `json:"lastId,omitempty"`
}

type ResultPaginated[T any] struct {
	Items     []T  `json:"items"`
	NextCursor int `json:"nextCursor"`
}