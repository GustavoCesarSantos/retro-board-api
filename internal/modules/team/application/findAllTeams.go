package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"
)

type IFindAllTeams interface {
    Execute(memberId int64) ([]*domain.Team, error)
}

type findAllTeams struct {
    repository db.ITeamRepository
}

func NewFindAllTeams(repository db.ITeamRepository) IFindAllTeams {
    return &findAllTeams{
        repository,
    }
}

func (fat *findAllTeams) Execute(memberId int64) ([]*domain.Team, error) {
    return fat.repository.FindAllByMemberId(memberId)
}
