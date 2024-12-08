package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
)

type ISaveOption interface {
    Execute(pollId int64, text string)
}

type saveOption struct {
    repository db.IOptionRepository
}

func NewSaveOption(repository db.IOptionRepository) ISaveOption {
    return &saveOption{
        repository,
    }
}

func (sp *saveOption) Execute(pollId int64, text string) {
    option := domain.NewOption(0, pollId, text)
    sp.repository.Save(*option)
}
