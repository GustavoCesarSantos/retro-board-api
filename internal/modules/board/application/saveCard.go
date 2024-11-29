package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
)

type ISaveCard interface {
    Execute(columnId int64, memberId int64, text string)
}

type saveCard struct {
    repository db.ICardRepository
}

func NewSaveCard(repository db.ICardRepository) ISaveCard {
    return &saveCard{
        repository,
    }
}

func (sc *saveCard) Execute(columnId int64, memberId int64, text string) {
    card := domain.NewCard(0, columnId, memberId, text)
    sc.repository.Save(*card)
}
