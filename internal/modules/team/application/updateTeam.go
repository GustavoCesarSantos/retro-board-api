package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"

type IUpdateTeam interface {
    Execute(teamId int64, name *string) error
}

type updateTeam struct {
    repository db.ITeamRepository
}

func NewUpdateTeam(repository db.ITeamRepository) IUpdateTeam {
    return &updateTeam{
        repository,
    }
}

func (up *updateTeam) Execute(teamId int64, name *string) error {
    team := struct {
         Name *string
    }{
        Name: name,
    }
    return up.repository.Update(teamId, team)
}
