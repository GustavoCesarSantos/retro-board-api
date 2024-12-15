package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"
)

type IFindTeam interface {
    Execute(teamId int64, memberId int64) (*domain.Team, error)
}

type findTeam struct {
    repository db.ITeamRepository
}

func NewFindTeam(repository db.ITeamRepository) IFindTeam {
    return &findTeam{
        repository,
    }
}

func (ft *findTeam) Execute(teamId int64, memberId int64) (*domain.Team, error) {
    return ft.repository.FindById(teamId, memberId)
}
