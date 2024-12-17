package domain

import "time"

type Poll struct {
    ID int64
	TeamId int64
    Name string
    CreatedAt time.Time
    UpdatedAt *time.Time
}

func NewPoll(
    id int64,
	teamId int64,
    name string,
) *Poll {
    return &Poll{
        ID: id,
		TeamId: teamId,
        Name: name,
        CreatedAt: time.Now(),
    }
}
