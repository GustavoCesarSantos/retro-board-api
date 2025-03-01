package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
)

type IUserRepository interface {
    FindByEmail(email string) (*domain.User, error)
    FindBySigninToken(signinToken string) (*domain.User, error)
	Save(user *domain.User) error
	UpdateSigninToken(userId int64, signinToken *domain.SigninToken) error
	UpdateVersion(user *domain.User) error
}
