package dtos

type CreateColumnRequest struct {
	Name   string       `json:"name"  example:"column1"`
	Color   string       `json:"color" example:"#FFFFFF"`
}