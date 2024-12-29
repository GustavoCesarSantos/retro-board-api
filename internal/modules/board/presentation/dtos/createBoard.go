package dtos

import "time"

type CreateBoardRequest struct {
	Name   string       `json:"name" example:"board1"`
}

type CreateBoardResponse struct {
	ID int64 `json:"id" example:"1"`
	Name string `json:"name" example:"board1"`
    CreatedAt time.Time `json:"created_at"`
}

func NewCreateBoardResponse(id int64, name string, createdAt time.Time) *CreateBoardResponse {
	return &CreateBoardResponse {
		ID: id,
		Name: name,
		CreatedAt: createdAt,
	}
}