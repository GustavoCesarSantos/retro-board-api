package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"

type IRemoveCard interface {
    Execute(cardId int64) error
}

type removeCard struct {
    repository db.ICardRepository
}

func NewRemoveCard(repository db.ICardRepository) IRemoveCard {
    return &removeCard{
        repository,
    }
}

func (rc *removeCard) Execute(cardId int64) error {
    return rc.repository.Delete(cardId)
}
