package dtos

import "time"

type ListAllTeamsResponse struct {
	ID int64 `json:"id" example:"1"`
	Name string `json:"name" example:"team1"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

type ListAllTeamsResponsePaginated struct {
	Items []*ListAllTeamsResponse `json:"items"`
	NextCursor int `json:"next_cursor" example:"0"`
}

func NewListAllTeamsResponse(
	id int64, 
	name string, 
	createdAt time.Time,
	updatedAt *time.Time,
) *ListAllTeamsResponse {
	return &ListAllTeamsResponse {
		ID: id,
		Name: name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}