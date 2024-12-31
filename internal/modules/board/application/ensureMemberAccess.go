package application

import (
	"errors"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
)

type IEnsureMemberAccess interface {
    Execute(teamId int64, memberId int64) error
}

type ensureMemberAccess struct {
    provider interfaces.ITeamMemberApi
}

func NewEnsureMemberAccess(provider interfaces.ITeamMemberApi) IEnsureMemberAccess {
    return &ensureMemberAccess{
        provider,
    }
}

func (ema *ensureMemberAccess) Execute(teamId int64, memberId int64) error {
    teamMembers, findErr := ema.provider.FindAllByTeamId(teamId)
    if findErr != nil {
        return findErr
    }
    found := false
	canEdit := false
    for _, teamMember := range teamMembers {
        if teamMember.MemberId == memberId  {
            found = true
			if teamMember.RoleId != 3 {
				canEdit = true
			} 
            break
        }
    }
    if !found {
        return errors.New("USER DOES NOT BELONG TO THE SPECIFIED TEAM")
    }
	if !canEdit {
		return errors.New("USER DOES NOT HAVE EDIT PERMISSION")
	}
    return nil
}
