package domain

import "time"

type Card struct {
    ID int64
	ColumnId int64
	MemberId int64
    Text string
    CreatedAt time.Time
    UpdatedAt *time.Time
}

func NewCard(
    id int64,
	columnId int64,
	memberId int64,
    text string,
) *Card {
    return &Card{
        ID: id,
		ColumnId: columnId,
		MemberId: memberId,
        Text: text,
        CreatedAt: time.Now(),
    }
}
