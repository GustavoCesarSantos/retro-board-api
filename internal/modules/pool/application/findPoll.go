package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
)

type IFindPoll interface {
    Execute(pollId int64) *domain.Poll
}

type findPoll struct {
    repository db.IPollRepository
}

func NewFindPoll(repository db.IPollRepository) IFindPoll {
    return &findPoll{
        repository,
    }
}

func (fp *findPoll) Execute(pollId int64) *domain.Poll {
    return fp.repository.FindById(pollId)
}
