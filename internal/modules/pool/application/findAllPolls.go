package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
)

type IFindAllPolls interface {
    Execute(teamId int64) []*domain.Poll
}

type findAllPolls struct {
    repository db.IPollRepository
}

func NewFindAllPolls(repository db.IPollRepository) IFindAllPolls {
    return &findAllPolls{
        repository,
    }
}

func (fap *findAllPolls) Execute(teamId int64) []*domain.Poll {
    return fap.repository.FindAllByTeamId(teamId)
}
