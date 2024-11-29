package domain

import "time"

type Board struct {
    ID int64
	TeamId int64
    Name string
	Active bool
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewBoard(
    id int64,
	teamId int64,
    name string,
) *Board {
    return &Board{
        ID: id,
		TeamId: teamId,
        Name: name,
        Active: true,
        CreatedAt: time.Now(),
    }
}
