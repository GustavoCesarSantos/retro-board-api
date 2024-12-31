package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/interfaces"

type IRemoveOption interface {
    Execute(optionId int64) error
}

type removeOption struct {
    repository db.IOptionRepository
}

func NewRemoveOption(repository db.IOptionRepository) IRemoveOption {
    return &removeOption{
        repository,
    }
}

func (ro *removeOption) Execute(optionId int64) error {
    return ro.repository.Delete(optionId)
}
