package dtos

import "time"

type ListAllBoardsResponse struct {
	ID int64 `json:"id" example:"1"`
	TeamId int64 `json:"team_id" example:"2"`
	Name string `json:"name" example:"board1"`
	Active bool `json:"active" example:"true"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

func NewListAllBoardsResponse(
	id int64,
	teamId int64,
	name string,
	active bool,
	createdAt time.Time,
	updatedAt *time.Time,
) *ListAllBoardsResponse {
	return &ListAllBoardsResponse {
		ID: id,
		TeamId: teamId,
		Name: name,
		Active: active,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}