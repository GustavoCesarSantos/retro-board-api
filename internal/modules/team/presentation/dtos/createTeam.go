package dtos

import "time"

type CreateTeamRequest struct {
	Name   string       `json:"name"`
}

type CreateTeamResponse struct {
	ID int64 `json:"id" example:"1"`
	Name string `json:"name" example:"team1"`
    CreatedAt time.Time `json:"created_at"`
}

func NewCreateTeamResponse(id int64, name string, createdAt time.Time) *CreateTeamResponse {
	return &CreateTeamResponse {
		ID: id,
		Name: name,
		CreatedAt: createdAt,
	}
}