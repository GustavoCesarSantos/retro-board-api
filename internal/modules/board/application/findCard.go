package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
)

type IFindCard interface {
    Execute(cardId int64) *domain.Card
}

type findCard struct {
    repository db.ICardRepository
}

func NewFindCard(repository db.ICardRepository) IFindCard {
    return &findCard{
        repository,
    }
}

func (fb *findCard) Execute(cardId int64) *domain.Card {
    return fb.repository.FindById(cardId)
}