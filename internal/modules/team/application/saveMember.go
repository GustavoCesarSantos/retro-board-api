package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"
)

type ISaveMember interface {
    Execute(teamId int64, memberId int64, roleId int64) error
}

type saveMember struct {
    repository db.ITeamMemberRepository
}

func NewSaveMember(repository db.ITeamMemberRepository) ISaveMember {
    return &saveMember{
        repository,
    }
}

func (sm *saveMember) Execute(teamId int64, memberId int64, roleId int64) error {
    teamMember := domain.NewTeamMember(0, teamId, memberId, roleId)
    return sm.repository.Save(teamMember)
}
