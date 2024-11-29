package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
)

type IMoveCardBetweenColumns interface {
    Execute(cardId int64, columnId int64)
}

type moveCardBetweenColumns struct {
    repository db.ICardRepository
}

func NewMoveCardBetweenColumns(repository db.ICardRepository) IMoveCardBetweenColumns {
    return &moveCardBetweenColumns{
        repository,
    }
}

func (mc *moveCardBetweenColumns) Execute(cardId int64, columnId int64) {
    mc.repository.MoveBetweenColumns(cardId, columnId)
}
