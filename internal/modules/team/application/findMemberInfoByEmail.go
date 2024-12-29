package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
)

type IFindMemberInfoByEmail interface {
    Execute(email string) (*domain.User, error)
}

type findMemberInfoByEmail struct {
    provider interfaces.IUserIdentityApi
}

func NewFindMemberInfoByEmail(provider interfaces.IUserIdentityApi) IFindMemberInfoByEmail {
    return &findMemberInfoByEmail{
        provider,
    }
}

func (fm *findMemberInfoByEmail) Execute(email string) (*domain.User, error) {
    return fm.provider.FindByEmail(email)
}
