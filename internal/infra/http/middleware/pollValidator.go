package middleware

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type PollValidator struct {
    provider interfaces.IPollApi
}

func NewPollValidator(provider interfaces.IPollApi) *PollValidator {
    return &PollValidator{
        provider,
    }
}


func (pv *PollValidator) EnsurePollOwnership(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metadataErr := utils.Envelope{
			"file": "pollValidator.go",
			"func": "pollValidator.EnsurePollOwnership",
			"line": 0,
		}
		teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
		if teamIdErr != nil {
			metadataErr["line"] = 30
			utils.BadRequestResponse(w, r, teamIdErr, metadataErr)
			return
		}
		pollId, pollIdErr := utils.ReadIDParam(r, "pollId")
		if pollIdErr != nil {
			metadataErr["line"] = 36
			utils.BadRequestResponse(w, r, pollIdErr, metadataErr)
			return
		}
		polls, findErr := pv.provider.FindAllPollsByTeamId(teamId)
		if findErr != nil {
			metadataErr["line"] = 42
			utils.NotFoundResponse(w, r, metadataErr)
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
			metadataErr["line"] = 54
			utils.ForbiddenResponse(w, r, utils.ErrPollNotInTeam, metadataErr)
			return 
		}
		next.ServeHTTP(w, r)
	})
}

func (pv *PollValidator) EnsureOptionOwnership(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metadataErr := utils.Envelope{
			"file": "pollValidator.go",
			"func": "pollValidator.EnsureOptionOwnership",
			"line": 0,
		}
		pollId, pollIdErr := utils.ReadIDParam(r, "pollId")
		if pollIdErr != nil {
			metadataErr["line"] = 71
			utils.BadRequestResponse(w, r, pollIdErr, metadataErr)
			return
		}
		optionId, optionIdErr := utils.ReadIDParam(r, "optionId")
		if optionIdErr != nil {
			metadataErr["line"] = 77
			utils.BadRequestResponse(w, r, optionIdErr, metadataErr)
			return
		}
		options, findErr := pv.provider.FindAllOptionsByPollId(pollId)
		if findErr != nil {
			metadataErr["line"] = 83
			utils.NotFoundResponse(w, r, metadataErr)
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
			metadataErr["line"] = 95
			utils.ForbiddenResponse(w, r, utils.ErrOptionNotInPoll, metadataErr)
			return 
		}
		next.ServeHTTP(w, r)
	})
}
