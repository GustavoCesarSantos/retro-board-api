package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/interfaces"
)


type IIncrementVersion interface {
    Execute(user *domain.User) error 
}

type incrementVersion struct {
    repository db.IUserRepository
}

func NewIncrementVersion(repository db.IUserRepository) IIncrementVersion {
    return &incrementVersion{
        repository,
    }
}

func (iv *incrementVersion) Execute(user *domain.User) error {
    return iv.repository.UpdateVersion(user)
}
