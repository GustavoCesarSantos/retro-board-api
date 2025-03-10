package db

import (
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type userRepository struct {
	users []domain.User
}

func NewUserRepository() db.IUserRepository {
	return &userRepository{
		users: []domain.User{},
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
    
func (ur *userRepository) FindBySigninToken(signinToken string) (*domain.User, error) {
    for _, user := range ur.users {
        if user.SigninToken.Token == signinToken && !user.SigninToken.ExpiresIn.Before(time.Now().UTC().Truncate(time.Second)) {
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

func (ur *userRepository) UpdateSigninToken(userId int64, signinToken *domain.SigninToken) error {
	for i := range ur.users {
		if ur.users[i].ID == userId {
			ur.users[i].SigninToken =  *signinToken
			break
		}
	}
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
