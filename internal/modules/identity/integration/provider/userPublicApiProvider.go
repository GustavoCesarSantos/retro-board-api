package provider

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
)

type userPublicApiProvider struct {
    repository db.IUserRepository
}

func NewUserPublicApiProvider(repository db.IUserRepository) interfaces.IUserIdentityApi {
    return &userPublicApiProvider{
        repository,
    }
}

func (upa userPublicApiProvider) FindByEmail(email string) (*domain.User, error) {
    return upa.repository.FindByEmail(email)
}
