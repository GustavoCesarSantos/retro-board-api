package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IFindCard interface {
    Execute(cardId int64) (*domain.Card, error)
}

type findCard struct {
    repository db.ICardRepository
}

func NewFindCard(repository db.ICardRepository) IFindCard {
    return &findCard{
        repository,
    }
}

func (fc *findCard) Execute(cardId int64) (*domain.Card, error) {
    return fc.repository.FindById(cardId)
}