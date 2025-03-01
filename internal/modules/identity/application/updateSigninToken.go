package application

import (
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/interfaces"
)

type IUpdateSigninToken interface {
    Execute(userId int64, token string) error
}

type updateSigninToken struct {
    repository db.IUserRepository
}

func NewUpdateSigninToken(repository db.IUserRepository) IUpdateSigninToken {
    return &updateSigninToken{
        repository,
    }
}

func (ust *updateSigninToken) Execute(userId int64, token string) error {
    signinToken := new(domain.SigninToken)
    signinToken.Token = token 
    signinToken.ExpiresIn = time.Now().Add(1 * time.Minute).UTC()
    err := ust.repository.UpdateSigninToken(userId, signinToken)
    return err
}
