package dtos

import "time"

type ListAllColumnsResponse struct {
	ID int64 `json:"id" example:"1"`
	BoardId int64 `json:"board_id" example:"2"`
	Name string `json:"name" example:"column1"`
	Color string `json:"color" example:"#FFFFFF"`
	Position int `json:"position" example:"1"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

type ListAllColumnsResponsePaginated struct {
	Items []*ListAllColumnsResponse `json:"items"`
	NextCursor int `json:"next_cursor" example:"0"`
}

func NewListAllColumnsResponse(
	id int64,
	boardId int64,
	name string,
	color string,
	position int,
	createdAt time.Time,
	updatedAt *time.Time,
) *ListAllColumnsResponse {
	return &ListAllColumnsResponse {
		ID: id,
		BoardId: boardId,
		Name: name,
		Color: color,
		Position: position,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}