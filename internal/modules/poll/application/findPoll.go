package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/interfaces"
)

type IFindPoll interface {
    Execute(pollId int64) (*domain.Poll, error)
}

type findPoll struct {
    repository db.IPollRepository
}

func NewFindPoll(repository db.IPollRepository) IFindPoll {
    return &findPoll{
        repository,
    }
}

func (fp *findPoll) Execute(pollId int64) (*domain.Poll, error) {
    return fp.repository.FindById(pollId)
}
