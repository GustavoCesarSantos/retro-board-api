package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type userRepository struct {
	users []domain.User
}

func NewUserRepository() db.IUserRepository {
	return &userRepository{
		users: []domain.User{
			*domain.NewUser(1, "Usuário 1", "usuario1@usuario.com"),
			*domain.NewUser(2, "Usuário 2", "usuario2@usuario.com"),
			*domain.NewUser(3, "Usuário 3", "usuario3@usuario.com"),
		},
	}
}

func (ur *userRepository) FindByEmail(email string) (*domain.User, error) {
    for _, user := range ur.users {
        if user.Email == email {
            return &user, nil
        }
    }
    return nil, utils.ErrRecordNotFound
}

func (ur *userRepository) Save(user *domain.User) error {
	user.ID = int64(len(ur.users) + 1)
	ur.users = append(ur.users, *user)
    return nil
}

func (ur *userRepository) UpdateVersion(user *domain.User) error {
	for i := range ur.users {
		if ur.users[i].ID == user.ID {
			ur.users[i].Version++
            user.Version = ur.users[i].Version 
			break
		}
	}
    return nil
}
