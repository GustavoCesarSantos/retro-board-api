package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"

type IPollRepository interface {
    FindAllByTeamId(teamId int64) ([]*domain.Poll, error)
    FindById(pollId int64) (*domain.Poll, error)
	Save(poll *domain.Poll) error
}