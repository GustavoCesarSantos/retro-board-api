package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/interfaces"
)

type ISaveVote interface {
    Execute(memberId int64, optionId int64) (*domain.Vote, error)
}

type saveVote struct {
    repository db.IVoteRepository
}

func NewSaveVote(repository db.IVoteRepository) ISaveVote {
    return &saveVote{
        repository,
    }
}

func (sp *saveVote) Execute(memberId int64, optionId int64) (*domain.Vote, error) {
    vote := domain.NewVote(0, memberId, optionId)
    err := sp.repository.Save(vote)
    return vote, err
}
