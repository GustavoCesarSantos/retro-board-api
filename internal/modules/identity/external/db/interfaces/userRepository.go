package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"

type IUserRepository interface {
    FindByEmail(email string) (*domain.User, error)
	Save(user *domain.User) error
	UpdateVersion(user *domain.User) error
}
