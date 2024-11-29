package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
)

type IUpdateCard interface {
    Execute(cardId int64, text *string)
}

type updateCard struct {
    repository db.ICardRepository
}

func NewUpdateCard(repository db.ICardRepository) IUpdateCard {
    return &updateCard{
        repository,
    }
}

func (uc *updateCard) Execute(cardId int64, text *string) {
    uc.repository.Update(cardId, text)
}
