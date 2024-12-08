package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
)

type ISaveVote interface {
    Execute(memberId int64, optionId int64)
}

type saveVote struct {
    repository db.IVoteRepository
}

func NewSaveVote(repository db.IVoteRepository) ISaveVote {
    return &saveVote{
        repository,
    }
}

func (sp *saveVote) Execute(memberId int64, optionId int64) {
    vote := domain.NewVote(0, memberId, optionId)
    sp.repository.Save(*vote)
}
