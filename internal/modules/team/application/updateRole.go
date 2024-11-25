package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/memory"
)

type IUpdateRole interface {
    Execute(teamId int64, memberId int64, role int64)
}

type updateRole struct {
    repository db.ITeamMemberRepository
}

func NewUpdateRole(repository db.ITeamMemberRepository) IUpdateRole {
    return &updateRole{
        repository,
    }
}

func (ur *updateRole) Execute(teamId int64, memberId int64, role int64) {
    ur.repository.UpdateRole(teamId, memberId, role)
}
