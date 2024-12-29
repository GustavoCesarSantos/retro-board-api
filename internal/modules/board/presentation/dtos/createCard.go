package dtos

import "time"

type CreateCardRequest struct {
	Text   string       `json:"text" example:"card1-text"`
}

type CreateCardResponse struct {
	ID int64 `json:"id" example:"1"`
	Text string `json:"text" example:"card1-text"`
    CreatedAt time.Time `json:"created_at"`
}

func NewCreateCardResponse(id int64, text string, createdAt time.Time) *CreateCardResponse {
	return &CreateCardResponse {
		ID: id,
		Text: text,
		CreatedAt: createdAt,
	}
}