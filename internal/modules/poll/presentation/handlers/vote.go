package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type Vote struct {
    notifySaveVote application.INotifySaveVote
	saveVote application.ISaveVote
}

func NewVote(
    notifySaveVote application.INotifySaveVote,
	saveVote application.ISaveVote,
) *Vote {
    return &Vote{
        notifySaveVote,
        saveVote,
    }
}

// @Summary      Vote in a poll
// @Description  Cast a vote for a specific option in a poll.
// @Tags         Poll
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int    true  "Team ID"
// @Param        pollId     path      int    true  "Poll ID"
// @Param        optionId   path      int    true  "Option ID"
// @Success      204        "Vote successfully recorded"
// @Failure      400        {object}  utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      403        {object}  utils.ErrorEnvelope  "User is not allowed to vote on this poll"
// @Failure      404        {object}  utils.ErrorEnvelope  "Poll or option not found"
// @Failure      500        {object}  utils.ErrorEnvelope  "Internal server error"
// @Router       /teams/:teamId/polls/:pollId/options/:optionId/vote [post]
func(v *Vote) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "vote.go",
		"func": "vote.Handle",
		"line": 0,
	}
	user := utils.ContextGetUser(r)
	pollId, pollIdErr := utils.ReadIDParam(r, "pollId")
	if pollIdErr != nil {
		utils.BadRequestResponse(w, r, pollIdErr, metadataErr)
		return
	}
	optionId, optionIdErr := utils.ReadIDParam(r, "optionId")
	if optionIdErr != nil {
		utils.BadRequestResponse(w, r, optionIdErr, metadataErr)
		return
	}
    _, saveErr := v.saveVote.Execute(user.ID, optionId)
	if saveErr != nil {
		utils.ServerErrorResponse(w, r, saveErr, metadataErr)
		return
	}
    v.notifySaveVote.Execute(pollId, optionId)
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
