package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/memory"
)

type ISaveUser interface {
    Execute(name string, email string)
}

type saveUser struct {
    repository db.IUserRepository
}

func NewSaveUser(repository db.IUserRepository) ISaveUser {
    return &saveUser{
        repository,
    }
}

func (su *saveUser) Execute(name string, email string) {
    user := domain.NewUser(0, name, email)
    su.repository.Save(*user)
}
