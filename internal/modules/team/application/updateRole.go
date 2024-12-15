package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"

type IUpdateRole interface {
    Execute(teamId int64, memberId int64, roleId int64) error
}

type updateRole struct {
    repository db.ITeamMemberRepository
}

func NewUpdateRole(repository db.ITeamMemberRepository) IUpdateRole {
    return &updateRole{
        repository,
    }
}

func (ur *updateRole) Execute(teamId int64, memberId int64, roleId int64) error {
    return ur.repository.UpdateRole(teamId, memberId, roleId)
}
