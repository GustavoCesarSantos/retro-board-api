package dtos

type MoveColumnToAnotherPositionRequest struct {
	NewPosition   int       `json:"new_position"  example:"3"`
}
