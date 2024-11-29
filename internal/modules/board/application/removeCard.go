package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"

type IRemoveCard interface {
    Execute(cardId int64)
}

type removeCard struct {
    repository db.ICardRepository
}

func NewRemoveCard(repository db.ICardRepository) IRemoveCard {
    return &removeCard{
        repository,
    }
}

func (rb *removeCard) Execute(cardId int64) {
    rb.repository.Delete(cardId)
}
