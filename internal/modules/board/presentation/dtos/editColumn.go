package dtos

type EditColumnRequest struct {
	Name   *string       `json:"name" example:"new-column1"`
	Color *string       `json:"color" example:"#000000"`
    Position *int `json:"position" example:"1"`
}
