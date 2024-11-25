package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/memory"
)

type IFindAllTeams interface {
    Execute(adminId int64) []*domain.Team
}

type findAllTeams struct {
    repository db.ITeamRepository
}

func NewFindAllTeams(repository db.ITeamRepository) IFindAllTeams {
    return &findAllTeams{
        repository,
    }
}

func (ft *findAllTeams) Execute(adminId int64) []*domain.Team {
    return ft.repository.FindAllByAdminId(adminId)
}
