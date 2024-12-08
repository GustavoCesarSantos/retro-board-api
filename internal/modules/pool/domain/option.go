package domain

import "time"

type Option struct {
    ID int64
	PollId int64
    Text string
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewOption(
    id int64,
	pollId int64,
    text string,
) *Option {
    return &Option{
        ID: id,
		PollId: pollId,
        Text: text,
        CreatedAt: time.Now(),
    }
}
