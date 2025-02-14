package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IMoveCardBetweenColumns interface {
    Execute(cardId int64, columnId int64) error
}

type moveCardBetweenColumns struct {
    repository db.ICardRepository
}

func NewMoveCardBetweenColumns(repository db.ICardRepository) IMoveCardBetweenColumns {
    return &moveCardBetweenColumns{
        repository,
    }
}

func (mc *moveCardBetweenColumns) Execute(cardId int64, columnId int64) error {
    card := struct{
        Text *string
        ColumnId *int64
    }{
        ColumnId: &columnId,
    }
    return mc.repository.Update(cardId, card)
}
