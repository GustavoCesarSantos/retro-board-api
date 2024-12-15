package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"

type IRemoveTeam interface {
    Execute(teamId int64, adminId int64) error
}

type removeTeam struct {
    repository db.ITeamRepository
}

func NewRemoveTeam(repository db.ITeamRepository) IRemoveTeam {
    return &removeTeam{
        repository,
    }
}

func (rt *removeTeam) Execute(teamId int64, adminId int64) error {
    return rt.repository.Delete(teamId, adminId)
}
