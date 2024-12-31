package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"

type IVoteRepository interface {
    CountByOptionId(optionId int64) (int, error)
	Save(vote *domain.Vote) error
}