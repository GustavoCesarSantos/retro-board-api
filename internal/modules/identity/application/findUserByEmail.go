package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/interfaces"
)

type IFindUserByEmail interface {
    Execute(email string) (*domain.User, error)
}

type findUserByEmail struct {
    repository db.IUserRepository
}

func NewFindUserByEmail(repository db.IUserRepository) IFindUserByEmail {
    return &findUserByEmail{
        repository,
    }
}

func (fu findUserByEmail) Execute(email string) (*domain.User, error) {
    return fu.repository.FindByEmail(email)
}
