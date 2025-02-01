package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type UpdateCardParams struct {
	Text *string
	ColumnId *int64
}

type ICardRepository interface {
	Delete(cardId int64) error
    FindAllByColumnId(columnId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Card], error)
	FindById(cardId int64) (*domain.Card, error)
	Save(card *domain.Card) error
	Update(cardId int64, card UpdateCardParams) error
}
