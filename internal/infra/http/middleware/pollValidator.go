package middleware

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type pollValidator struct {
    provider interfaces.IPollApi
}

func NewPollValidator(provider interfaces.IPollApi) *pollValidator {
    return &pollValidator{
        provider,
    }
}


func (pv *pollValidator) EnsurePollOwnership(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
		if teamIdErr != nil {
			utils.BadRequestResponse(w, r, teamIdErr)
			return
		}
		pollId, pollIdErr := utils.ReadIDParam(r, "pollId")
		if pollIdErr != nil {
			utils.BadRequestResponse(w, r, pollIdErr)
			return
		}
		polls, findErr := pv.provider.FindAllPollsByTeamId(teamId)
		if findErr != nil {
			utils.NotFoundResponse(w, r)
			return 
		}
		found := false
		for _, poll := range polls {
			if poll.ID == pollId {
				found = true
				break
			}
		}
		if !found {
			utils.ForbiddenResponse(w, r, utils.ErrPollNotInTeam)
			return 
		}
		next.ServeHTTP(w, r)
	})
}

func (pv *pollValidator) EnsureOptionOwnership(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pollId, pollIdErr := utils.ReadIDParam(r, "pollId")
		if pollIdErr != nil {
			utils.BadRequestResponse(w, r, pollIdErr)
			return
		}
		optionId, optionIdErr := utils.ReadIDParam(r, "optionId")
		if optionIdErr != nil {
			utils.BadRequestResponse(w, r, optionIdErr)
			return
		}
		options, findErr := pv.provider.FindAllOptionsByPollId(pollId)
		if findErr != nil {
			utils.NotFoundResponse(w, r)
			return 
		}
		found := false
		for _, option := range options {
			if option.ID == optionId {
				found = true
				break
			}
		}
		if !found {
			utils.ForbiddenResponse(w, r, utils.ErrOptionNotInPoll)
			return 
		}
		next.ServeHTTP(w, r)
	})
}
