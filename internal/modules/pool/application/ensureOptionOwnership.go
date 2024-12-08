package application

import (
	"errors"

	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
)

type IEnsureOptionOwnership interface {
    Execute(pollId int64, optionId int64) error
}

type ensureOptionOwnership struct {
    repository db.IOptionRepository
}

func NewEnsureOptionOwnership(repository db.IOptionRepository) IEnsureOptionOwnership {
    return &ensureOptionOwnership{
        repository,
    }
}

func (epo *ensureOptionOwnership) Execute(pollId int64, optionId int64) error {
    options := epo.repository.FindAllByPollId(pollId)
    found := false
    for _, option := range options {
        if option.ID == optionId {
            found = true
            break
        }
    }
    if !found {
        return errors.New("OPTION DOES NOT BELONG TO THE SPECIFIED POLL")
    }
    return nil
}
