package dtos

type MoveCardtoAnotherColumnRequest struct {
	NewColumnId   int64       `json:"new_column_id"  example:"3"`
}