package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/memory"
)

type ISaveTeam interface {
    Execute(name string, adminId int64)
}

type saveTeam struct {
    repository db.ITeamRepository
}

func NewSaveTeam(repository db.ITeamRepository) ISaveTeam {
    return &saveTeam{
        repository,
    }
}

func (su *saveTeam) Execute(name string, adminId int64) {
    team := domain.NewTeam(0, name, adminId)
    su.repository.Save(*team)
}
