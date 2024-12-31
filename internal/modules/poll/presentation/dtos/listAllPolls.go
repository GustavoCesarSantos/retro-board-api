package dtos

import "time"

type ListAllPollsResponse struct {
	ID int64 `json:"id" example:"1"`
	TeamId int64 `json:"team_id" example:"2"`
	Name string `json:"name" example:"poll1"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

func NewListAllPollsResponse(
	id int64,
	teamId int64,
	name string,
	createdAt time.Time,
	updatedAt *time.Time,
) *ListAllPollsResponse {
	return &ListAllPollsResponse {
		ID: id,
		TeamId: teamId,
		Name: name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}