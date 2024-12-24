package dtos

import "time"

type ListCardResponse struct {
	ID int64 `json:"id" example:"1"`
	ColumnId int64 `json:"column_id" example:"2"`
	MemberId int64 `json:"member_id" example:"2"`
	Text string `json:"text" example:"card1-text"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

func NewListCardResponse(
	id int64,
	columnId int64,
	memberId int64,
	text string,
	createdAt time.Time,
	updatedAt *time.Time,
) *ListCardResponse {
	return &ListCardResponse {
		ID: id,
		ColumnId: columnId,
		MemberId: memberId,
		Text: text,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}