package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"

type ICardRepository interface {
	Delete(cardId int64)
    FindAllByColumnId(columnId int64) []*domain.Card
	FindById(cardId int64) *domain.Card
	MoveBetweenColumns(cardId int64, columnId int64)
	Save(card domain.Card)
	Update(cardId int64, text *string)
}

type cardRepository struct {
	cards []domain.Card
}

func NewCardRepository() ICardRepository {
	return &cardRepository{
		cards: []domain.Card{
			*domain.NewCard(1, 1, 1, "Texto do card 1"),
			*domain.NewCard(2, 1, 1, "Texto do card 2"),
			*domain.NewCard(3, 1, 1, "Texto do card 3"),
		},
	}
}

func (cr *cardRepository) Delete(cardId int64) {
    i := 0
	for _, card := range cr.cards {
		if !(card.ID == cardId) {
			cr.cards[i] = card
			i++
		}
	}
	cr.cards = cr.cards[:i]
}

func (cr *cardRepository) FindAllByColumnId(columnId int64) []*domain.Card {
    var cards []*domain.Card
    for _, card := range cr.cards {
        if card.ColumnId == columnId {
            cards = append(cards, &card)
        }
    }
    return cards
}

func (cr *cardRepository) FindById(cardId int64) *domain.Card {
    for _, card := range cr.cards {
        if card.ID == cardId {
            return &card
        }
    }
    return nil
}

func (cr *cardRepository) MoveBetweenColumns(cardId int64, columnId int64) {
    
}

func (cr *cardRepository) Save(card domain.Card) {
	cr.cards = append(cr.cards, card)
}

func (cr *cardRepository) Update(cardId int64, text *string) {
    for i := range cr.cards {
		if cr.cards[i].ID == cardId {
			if text != nil {
				cr.cards[i].Text = *text
			}
			break
		}
	}
}
