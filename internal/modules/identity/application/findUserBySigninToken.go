package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/interfaces"
)

type IFindUserBySigninToken interface {
    Execute(signinToken string) (*domain.User, error)
}

type findUserBySigninToken struct {
    repository db.IUserRepository
}

func NewFindUserBySigninToken(repository db.IUserRepository) IFindUserBySigninToken {
    return &findUserBySigninToken{
        repository,
    }
}

func (fu findUserBySigninToken) Execute(signinToken string) (*domain.User, error) {
    return fu.repository.FindBySigninToken(signinToken)
}
