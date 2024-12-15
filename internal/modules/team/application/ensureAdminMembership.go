package application

import (
	"errors"

	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"
)

type IEnsureAdminMembership interface {
    Execute(teamId int64, adminId int64) error
}

type ensureAdminMembership struct {
    repository db.ITeamMemberRepository
}

func NewEnsureAdminMembership(repository db.ITeamMemberRepository) IEnsureAdminMembership {
    return &ensureAdminMembership{
        repository,
    }
}

func (eam *ensureAdminMembership) Execute(teamId int64, adminId int64) error {
    _, err := eam.repository.FindTeamAdminByMemberId(teamId, adminId)
	if err != nil {
		return errors.New("USER DOES NOT HAVE ADMIN PERMISSIONS")
	}
	return nil
}
