package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IUpdateCard interface {
    Execute(cardId int64, text *string) error
}

type updateCard struct {
    repository db.ICardRepository
}

func NewUpdateCard(repository db.ICardRepository) IUpdateCard {
    return &updateCard{
        repository,
    }
}

func (uc *updateCard) Execute(cardId int64, text *string) error {
    card := struct{
        Text *string
        ColumnId *int64
    }{
        Text: text,
        ColumnId: nil,
    }
    return uc.repository.Update(cardId, card)
}
