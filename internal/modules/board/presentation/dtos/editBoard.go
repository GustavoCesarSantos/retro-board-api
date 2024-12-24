package dtos

type EditBoardRequest struct {
	Name   *string       `json:"name" example:"new-board1"`
	Active *bool       `json:"active" example:"true"`
}