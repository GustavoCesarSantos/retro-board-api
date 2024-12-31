package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/interfaces"
)

type ISavePoll interface {
    Execute(teamId int64, name string) (*domain.Poll, error)
}

type savePoll struct {
    repository db.IPollRepository
}

func NewSavePoll(repository db.IPollRepository) ISavePoll {
    return &savePoll{
        repository,
    }
}

func (sp *savePoll) Execute(teamId int64, name string) (*domain.Poll, error) {
    poll := domain.NewPoll(0, teamId, name)
    err := sp.repository.Save(poll)
    return poll, err
}
