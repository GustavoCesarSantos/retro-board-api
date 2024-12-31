package domain

import "time"

type Vote struct {
    ID int64
	MemberId int64
	OptionId int64
    CreatedAt time.Time
}

func NewVote(
    id int64,
	memberId int64,
	optionId int64,
) *Vote {
    return &Vote{
        ID: id,
		MemberId: memberId,
        OptionId: optionId,
        CreatedAt: time.Now(),
    }
}
