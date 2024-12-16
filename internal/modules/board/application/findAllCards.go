package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IFindAllCards interface {
    Execute(columnId int64) ([]*domain.Card, error)
}

type findAllCards struct {
    repository db.ICardRepository
}

func NewFindAllCards(repository db.ICardRepository) IFindAllCards {
    return &findAllCards{
        repository,
    }
}

func (fac *findAllCards) Execute(columnId int64) ([]*domain.Card, error) {
    return fac.repository.FindAllByColumnId(columnId)
}
