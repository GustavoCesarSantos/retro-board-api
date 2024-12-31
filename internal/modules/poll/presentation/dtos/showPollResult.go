package dtos

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/application"


type ShowPollResultResponse application.CountVotesResult

func NewShowPollResultResponse(
	options map[int64]application.Option,
	winner  []application.Winner,
	total int,
) *ShowPollResultResponse {
	return &ShowPollResultResponse {
		Options: options,
		Winner: winner,
		Total: total,
	}
}