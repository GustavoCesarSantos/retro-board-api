package interfaces

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"

type IUserIdentityApi interface {
	FindByEmail(email string) (*domain.User, error)
}