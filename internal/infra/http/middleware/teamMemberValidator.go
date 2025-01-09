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
		metadataErr := utils.Envelope{
			"file": "teamMemberValidator.go",
			"func": "teamMemberValidator.EnsureMemberAccess",
			"line": 0,
		}
		user := utils.ContextGetUser(r)
		teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
		if teamIdErr != nil {
			metadataErr["line"] = 30
			utils.BadRequestResponse(w, r, teamIdErr, metadataErr)
			return
		}
		teamMembers, findErr := tmv.provider.FindAllByTeamId(teamId)
		if findErr != nil {
			metadataErr["line"] = 36
			utils.NotFoundResponse(w, r, metadataErr)
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
			metadataErr["line"] = 52
			utils.ForbiddenResponse(w, r, utils.ErrUserNotInTeam, metadataErr)
			return 
		}
		if !canEdit {
			metadataErr["line"] = 57
			utils.ForbiddenResponse(w, r, utils.ErrUserNoEditPermission, metadataErr)
			return
		}
		r = utils.ContextSetUser(r, user)
		next.ServeHTTP(w, r)
	})
}
