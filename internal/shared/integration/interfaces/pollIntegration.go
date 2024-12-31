package interfaces

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"

type IPollApi interface {
	FindAllOptionsByPollId(pollId int64) ([]*domain.Option, error)
	FindAllPollsByTeamId(teamId int64) ([]*domain.Poll, error)
}