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

func NewUserRepository(db *sql.DB) db.IUserRepository {
	return &userRepository{
        DB: db,
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

func (ur *userRepository) FindBySigninToken(signinToken string) (*domain.User, error) {
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
            signin_token ->>'token' = $1
            AND (signin_token ->>'expires_in')::timestamp >= NOW()
    `
	var user domain.User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := ur.DB.QueryRowContext(ctx, query, signinToken).Scan(
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
            email
        )
        VALUES (
            $1,
            $2
        )
        RETURNING
            id,
            version,
            created_at
    `
	args := []any{user.Name, user.Email}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    return ur.DB.QueryRowContext(ctx, query, args...).Scan(
        &user.ID,
        &user.Version,
        &user.CreatedAt,
    )
}

func (ur *userRepository) UpdateSigninToken(userId int64, signinToken *domain.SigninToken) error {
    query := `
        UPDATE
            users 
        SET
            signin_token = $1
        WHERE
            id = $2
    `
	args := []any{
        signinToken,
        userId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    result, err := ur.DB.ExecContext(ctx, query, args...)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return utils.ErrRecordNotFound
		default:
			return err

		}
	}
    rows, err := result.RowsAffected()
    if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return utils.ErrRecordNotFound
		default:
			return err

		}
	}
    if rows == 0 {
        return utils.ErrRecordNotFound
    }
	return nil
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
