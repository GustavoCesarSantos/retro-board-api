package application

import (
	"errors"

	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IEnsureCardOwnership interface {
    Execute(columnId int64, cardId int64) error
}

type ensureCardOwnership struct {
    repository db.ICardRepository
}

func NewEnsureCardOwnership(repository db.ICardRepository) IEnsureCardOwnership {
    return &ensureCardOwnership{
        repository,
    }
}

func (eco *ensureCardOwnership) Execute(columnId int64, cardId int64) error {
    cards, findErr := eco.repository.FindAllByColumnId(columnId)
    if findErr != nil {
        return findErr
    }
    found := false
    for _, card := range cards {
        if card.ID == cardId {
            found = true
            break
        }
    }
    if !found {
        return errors.New("CARD DOES NOT BELONG TO THE SPECIFIED COLUMN")
    }
    return nil
}
