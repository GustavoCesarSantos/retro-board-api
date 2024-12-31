package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/interfaces"
)

type ISaveOption interface {
    Execute(pollId int64, text string) (*domain.Option, error)
}

type saveOption struct {
    repository db.IOptionRepository
}

func NewSaveOption(repository db.IOptionRepository) ISaveOption {
    return &saveOption{
        repository,
    }
}

func (sp *saveOption) Execute(pollId int64, text string) (*domain.Option, error) {
    option := domain.NewOption(0, pollId, text)
    err := sp.repository.Save(option)
    return option, err
}
