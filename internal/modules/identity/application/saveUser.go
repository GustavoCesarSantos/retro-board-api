package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/interfaces"
)

type ISaveUser interface {
    Execute(name string, email string) (*domain.User, error)
}

type saveUser struct {
    repository db.IUserRepository
}

func NewSaveUser(repository db.IUserRepository) ISaveUser {
    return &saveUser{
        repository,
    }
}

func (su *saveUser) Execute(name string, email string) (*domain.User, error) {
    user := domain.NewUser(name, email)
    err := su.repository.Save(user)
    return user, err
}
