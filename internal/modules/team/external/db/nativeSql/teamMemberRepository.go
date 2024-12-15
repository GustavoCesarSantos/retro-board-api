package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type teamMemberRepository struct {
    DB *sql.DB
}

func NewTeamMemberRepository(db *sql.DB) db.ITeamMemberRepository {
	return &teamMemberRepository{
        DB:db,
	}
}

func (tm *teamMemberRepository) Delete(teamId int64, memberId int64,) error {
    query := `
        DELETE FROM
            teamMembers
        WHERE
            teamId = $1
            AND memberId = $2;
    `
	args := []any{teamId, memberId}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := tm.DB.ExecContext(ctx, query, args...)
    if err != nil {
        return err
    }
    rowsAffected, rowsAffectedErr := result.RowsAffected()
    if rowsAffectedErr != nil {
        return rowsAffectedErr
    }
    if rowsAffected == 0 {
        return utils.ErrRecordNotFound
    }
    return nil
}

func (tm *teamMemberRepository) FindTeamAdminByMemberId(teamId int64, memberId int64) (*domain.TeamMember, error) {
    query := `
        SELECT
            id,
            created_at,
            updated_at
        FROM
            teamMembers
        WHERE
            memberId = $1
            AND roleId = 1
            AND teamId = $2;
    `
    args := []any{memberId, teamId}
	var teamMember domain.TeamMember
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := tm.DB.QueryRowContext(ctx, query, args...).Scan(
		&teamMember.ID,
        &teamMember.CreatedAt,
		&teamMember.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, utils.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &teamMember, nil
}

func (tm *teamMemberRepository) Save(teamMember *domain.TeamMember) error {
    query := `
        INSERT INTO teamMembers (
            memberId,
            teamId,
            roleId
        )
        VALUES (
            $1,
            $2,
            $3
        )
        RETURNING
            id,
            created_at
    `
	args := []any{teamMember.MemberId, teamMember.TeamId, teamMember.RoleId}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    return tm.DB.QueryRowContext(ctx, query, args...).Scan(
        &teamMember.ID,
        &teamMember.CreatedAt,
    )
}

func (tm *teamMemberRepository) UpdateRole(teamId int64, memberId int64, roleId int64) error {
    query := `
        UPDATE
            teamMembers 
        SET
            roleId = $1
        WHERE
            teamId = $2
            AND memberId = $3;
    `
	args := []any{
        roleId,
        teamId,
        memberId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := tm.DB.ExecContext(ctx, query, args...)
    if err != nil {
        return err
    }
    rowsAffected, rowsAffectedErr := result.RowsAffected()
    if rowsAffectedErr != nil {
        return rowsAffectedErr
    }
    if rowsAffected == 0 {
        return utils.ErrRecordNotFound
    }
    return nil
}
