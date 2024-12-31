package middleware

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type teamMemberValidator struct {
    provider interfaces.ITeamMemberApi
}

func NewTeamMemberValidator(provider interfaces.ITeamMemberApi) *teamMemberValidator {
    return &teamMemberValidator{
        provider,
    }
}

func (tmv *teamMemberValidator) EnsureMemberAccess(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := utils.ContextGetUser(r)
		teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
		if teamIdErr != nil {
			utils.BadRequestResponse(w, r, teamIdErr)
			return
		}
		teamMembers, findErr := tmv.provider.FindAllByTeamId(teamId)
		if findErr != nil {
			utils.NotFoundResponse(w, r)
			return
		}
		found := false
		canEdit := false
		for _, teamMember := range teamMembers {
			if teamMember.MemberId == user.ID  {
				found = true
				if teamMember.RoleId != 3 {
					canEdit = true
				} 
				break
			}
		}
		if !found {
			utils.ForbiddenResponse(w, r, utils.ErrUserNotInTeam)
			return 
		}
		if !canEdit {
			utils.ForbiddenResponse(w, r, utils.ErrUserNoEditPermission)
			return
		}
		r = utils.ContextSetUser(r, user)
		next.ServeHTTP(w, r)
	})
}
