package domain

import "time"

type Column struct {
    ID int64
	BoardId int64
    Name string
    Color string
    Position int
    CreatedAt time.Time
    UpdatedAt *time.Time
}

func NewColumn(
    id int64,
	boardId int64,
    name string,
	color string,
    position int,
) *Column {
    return &Column{
        ID: id,
		BoardId: boardId,
        Name: name,
		Color: color,
        Position: position,
        CreatedAt: time.Now(),
    }
}
