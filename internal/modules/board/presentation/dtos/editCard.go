package dtos

type EditCardRequest struct {
	Text   *string       `json:"text" example:"new-card1-text"`
}