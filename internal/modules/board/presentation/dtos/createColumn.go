package dtos

import "time"

type CreateColumnRequest struct {
	Name   string       `json:"name"  example:"column1"`
	Color   string       `json:"color" example:"#FFFFFF"`
}

type CreateColumnResponse struct {
	ID int64 `json:"id" example:"1"`
	Name string `json:"name" example:"column1"`
	Color   string       `json:"color" example:"#FFFFFF"`
    CreatedAt time.Time `json:"created_at"`
}

func NewCreateColumnResponse(id int64, name string, color string, createdAt time.Time) *CreateColumnResponse {
	return &CreateColumnResponse {
		ID: id,
		Name: name,
		Color: color,
		CreatedAt: createdAt,
	}
}