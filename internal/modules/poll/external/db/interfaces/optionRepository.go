package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"

type IOptionRepository interface {
	Delete(optionId int64) error
    FindAllByPollId(pollId int64) ([]*domain.Option, error)
	Save(option *domain.Option) error
}