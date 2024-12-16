package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)


type cardRepository struct {
	cards []domain.Card
}

func NewCardRepository() db.ICardRepository {
	return &cardRepository{
		cards: []domain.Card{
			*domain.NewCard(1, 1, 1, "Texto do card 1"),
			*domain.NewCard(2, 1, 1, "Texto do card 2"),
			*domain.NewCard(3, 1, 1, "Texto do card 3"),
		},
	}
}

func (cr *cardRepository) Delete(cardId int64) error {
    i := 0
	for _, card := range cr.cards {
		if !(card.ID == cardId) {
			cr.cards[i] = card
			i++
		}
	}
	cr.cards = cr.cards[:i]
    return nil
}

func (cr *cardRepository) FindAllByColumnId(columnId int64) ([]*domain.Card, error) {
    var cards []*domain.Card
    for _, card := range cr.cards {
        if card.ColumnId == columnId {
            cards = append(cards, &card)
        }
    }
    return cards, nil
}

func (cr *cardRepository) FindById(cardId int64) (*domain.Card, error) {
    for _, card := range cr.cards {
        if card.ID == cardId {
            return &card, nil
        }
    }
    return nil, utils.ErrRecordNotFound
}

func (cr *cardRepository) MoveBetweenColumns(cardId int64, columnId int64) error {
   return nil 
}

func (cr *cardRepository) Save(card *domain.Card) error {
	cr.cards = append(cr.cards, *card)
    return nil
}

func (cr *cardRepository) Update(cardId int64, card db.UpdateCardParams) error {
    for i := range cr.cards {
		if cr.cards[i].ID == cardId {
			if card.Text != nil {
				cr.cards[i].Text = *card.Text
			}
			if card.ColumnId != nil {
				cr.cards[i].ColumnId = *card.ColumnId
			}
			break
		}
	}
    return nil
}
