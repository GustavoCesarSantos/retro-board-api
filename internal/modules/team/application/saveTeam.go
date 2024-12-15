package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"
)

type ISaveTeam interface {
    Execute(name string, adminId int64) (*domain.Team, error)
}

type saveTeam struct {
    repository db.ITeamRepository
}

func NewSaveTeam(repository db.ITeamRepository) ISaveTeam {
    return &saveTeam{
        repository,
    }
}

func (st *saveTeam) Execute(name string, adminId int64) (*domain.Team, error) {
    team := domain.NewTeam(0, name, adminId)
    err := st.repository.Save(team)
    return team, err
}
