package poll

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listPoll struct {
	findPoll application.IFindPoll
}

func NewListPoll(
	findPoll application.IFindPoll,
) *listPoll {
    return &listPoll {
		findPoll,
    }
}

type ListPollEnvelop struct {
	Poll *dtos.ListPollResponse `json:"poll"`
}

// @Summary      Get poll details
// @Description  Retrieves detailed information about a specific poll.
// @Tags         Poll
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int    true  "Team ID"
// @Param        pollId     path      int    true  "Poll ID"
// @Success      200        {object}  poll.ListPollEnvelop  "Poll details"
// @Failure      400        {object}  utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404        {object}  utils.ErrorEnvelope  "Poll not found"
// @Failure      500        {object}  utils.ErrorEnvelope  "Internal server error"
// @Router       /teams/:teamId/polls/:pollId [get]
func(lp *listPoll) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "listPoll.go",
		"func": "listPoll.Handle",
		"line": 0,
	}
	pollId, pollIdErr := utils.ReadIDParam(r, "pollId")
	if pollIdErr != nil {
		utils.BadRequestResponse(w, r, pollIdErr, metadataErr)
		return
	}
	poll, findErr := lp.findPoll.Execute(pollId)
	if findErr != nil {
		switch {
		case errors.Is(findErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r, metadataErr)
		default:
            utils.ServerErrorResponse(w, r, findErr, metadataErr)
		}
		return
    }
	response := dtos.NewListPollResponse(
		poll.ID,
		poll.TeamId,
		poll.Name,
		poll.CreatedAt,
		poll.UpdatedAt,
	)
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"poll": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
