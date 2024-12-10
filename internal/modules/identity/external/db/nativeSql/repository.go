package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type userRepository struct {
    DB *sql.DB
}

func NewUserRepository(DB *sql.DB) db.IUserRepository {
	return &userRepository{
        DB,
	}
}

func (ur *userRepository) FindByEmail(email string) (*domain.User, error) {
    query := `
        SELECT
            id,
            name,
            email,
            version,
            created_at,
            updated_at
        FROM
            users
        WHERE
            email = $1
    `
	var user domain.User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := ur.DB.QueryRowContext(ctx, query, email).Scan(
		&[]byte{},
		&user.ID,
        &user.Name,
        &user.Email,
        &user.Version,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, utils.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (ur *userRepository) Save(user *domain.User) error {
    query := `
        INSERT INTO users (
            name,
            email,
            version
        )
        VALUES (
            $1,
            $2,
            $3
        )
        RETURNING
            id,
            version,
            created_at
    `
	args := []any{user.Name, user.Email, user.Version}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    return ur.DB.QueryRowContext(ctx, query, args...).Scan(
        &user.ID,
        &user.Version,
        &user.CreatedAt,
    )
}

func (ur *userRepository) UpdateVersion(user *domain.User) error {
    query := `
        UPDATE
            users 
        SET
            version = version + 1
        WHERE
            id = $1
        RETURNING version
    `
	args := []any{
        user.ID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    err := ur.DB.QueryRowContext(ctx, query, args...).Scan(&user.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return utils.ErrRecordNotFound
		default:
			return err

		}
	}
	return nil
}
