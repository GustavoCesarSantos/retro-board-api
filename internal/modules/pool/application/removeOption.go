package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"

type IRemoveOption interface {
    Execute(optionId int64)
}

type removeOption struct {
    repository db.IOptionRepository
}

func NewRemoveOption(repository db.IOptionRepository) IRemoveOption {
    return &removeOption{
        repository,
    }
}

func (ro *removeOption) Execute(optionId int64) {
    ro.repository.Delete(optionId)
}
