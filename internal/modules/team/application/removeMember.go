package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/memory"
)

type IRemoveMember interface {
    Execute(teamId int64, memberId int64)
}

type removeMember struct {
    repository db.ITeamMemberRepository
}

func NewRemoveMember(repository db.ITeamMemberRepository) IRemoveMember {
    return &removeMember{
        repository,
    }
}

func (rm *removeMember) Execute(teamId int64, memberId int64) {
    rm.repository.Delete(teamId	, memberId)
}
