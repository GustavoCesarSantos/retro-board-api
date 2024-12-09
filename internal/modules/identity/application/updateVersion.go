package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/memory"

type IUpdateVersion interface {
    Execute(userId int64, version int)
}

type updateVersion struct {
    repository db.IUserRepository
}

func NewUpdateVersion(repository db.IUserRepository) IUpdateVersion {
    return &updateVersion{
        repository,
    }
}

func (ur *updateVersion) Execute(userId int64, version int) {
    ur.repository.UpdateVersion(userId, version)
}
