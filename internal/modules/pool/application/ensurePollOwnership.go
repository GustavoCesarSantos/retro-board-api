package application

import (
	"errors"

	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
)

type IEnsurePollOwnership interface {
    Execute(teamId int64, pollId int64) error
}

type ensurePollOwnership struct {
    repository db.IPollRepository
}

func NewEnsurePollOwnership(repository db.IPollRepository) IEnsurePollOwnership {
    return &ensurePollOwnership{
        repository,
    }
}

func (epo *ensurePollOwnership) Execute(teamId int64, pollId int64) error {
    polls := epo.repository.FindAllByTeamId(teamId)
    found := false
    for _, poll := range polls {
        if poll.ID == pollId {
            found = true
            break
        }
    }
    if !found {
        return errors.New("POLL DOES NOT BELONG TO THE SPECIFIED TEAM")
    }
    return nil
}
