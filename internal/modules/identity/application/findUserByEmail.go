package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/memory"
)

type IFindUserByEmail interface {
    Execute(email string) *domain.User
}

type findUserByEmail struct {
    repository db.IUserRepository
}

func NewFindUserByEmail(repository db.IUserRepository) IFindUserByEmail {
    return &findUserByEmail{
        repository,
    }
}

func (fu findUserByEmail) Execute(email string) *domain.User {
    return fu.repository.FindByEmail(email)
}
