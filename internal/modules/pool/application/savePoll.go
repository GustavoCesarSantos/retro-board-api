package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
)

type ISavePoll interface {
    Execute(teamId int64, name string) int64
}

type savePoll struct {
    repository db.IPollRepository
}

func NewSavePoll(repository db.IPollRepository) ISavePoll {
    return &savePoll{
        repository,
    }
}

func (sp *savePoll) Execute(teamId int64, name string) int64 {
    poll := domain.NewPoll(0, teamId, name)
    sp.repository.Save(*poll)
    return poll.ID
}
