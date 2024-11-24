package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"

type IUserRepository interface {
	Save(user domain.User)
    FindByEmail(email string) *domain.User
}

type userRepository struct {
	users []domain.User
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		users: []domain.User{
			*domain.NewUser(1, "Usuário 1", "usuario1@usuario.com"),
			*domain.NewUser(2, "Usuário 2", "usuario2@usuario.com"),
			*domain.NewUser(3, "Usuário 3", "usuario3@usuario.com"),
		},
	}
}

func (ur *userRepository) Save(user domain.User) {
	ur.users = append(ur.users, user)
}

func (ur *userRepository) FindByEmail(email string) *domain.User {
    for _, user := range ur.users {
        if user.Email == email {
            return &user
        }
    }
    return nil
}
