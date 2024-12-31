package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/interfaces"
)

type IFindAllPolls interface {
    Execute(teamId int64) ([]*domain.Poll, error)
}

type findAllPolls struct {
    repository db.IPollRepository
}

func NewFindAllPolls(repository db.IPollRepository) IFindAllPolls {
    return &findAllPolls{
        repository,
    }
}

func (fap *findAllPolls) Execute(teamId int64) ([]*domain.Poll, error) {
    return fap.repository.FindAllByTeamId(teamId)
}
