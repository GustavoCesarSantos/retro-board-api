package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/memory"
)

type ISaveMember interface {
    Execute(teamId int64, memberId int64, role int64)
}

type saveMember struct {
    repository db.ITeamMemberRepository
}

func NewSaveMember(repository db.ITeamMemberRepository) ISaveMember {
    return &saveMember{
        repository,
    }
}

func (sm *saveMember) Execute(teamId int64, memberId int64, role int64) {
    teamMember := domain.NewTeamMember(0, teamId, memberId, role)
    sm.repository.Save(*teamMember)
}
