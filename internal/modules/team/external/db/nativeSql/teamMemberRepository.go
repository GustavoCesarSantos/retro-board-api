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
		DB: db,
	}
}

func (tm *teamMemberRepository) Delete(teamId int64, memberId int64) error {
	query := `
        DELETE FROM
            team_members
        WHERE
            team_id = $1
            AND member_id = $2;
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

func (tm *teamMemberRepository) FindAllByTeamId(teamId int64) ([]*domain.TeamMember, error) {
	query := `
        SELECT
            id,
            team_id,
            member_id,
            role_id,
            status,
            created_at,
            updated_at
        FROM
            team_members
        WHERE
            team_id = $1
            AND status = 'active';
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := tm.DB.QueryContext(ctx, query, teamId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	teamMembers := []*domain.TeamMember{}
	for rows.Next() {
		var teamMember domain.TeamMember
		err := rows.Scan(
			&teamMember.ID,
			&teamMember.TeamId,
			&teamMember.MemberId,
			&teamMember.RoleId,
            &teamMember.Status,
			&teamMember.CreatedAt,
			&teamMember.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		teamMembers = append(teamMembers, &teamMember)
	}
	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}
	return teamMembers, nil
}

func (tm *teamMemberRepository) FindTeamAdminByMemberId(
	teamId int64,
	memberId int64,
) (*domain.TeamMember, error) {
	query := `
        SELECT
            id,
            created_at,
            updated_at
        FROM
            team_members
        WHERE
            member_id = $1
            AND role_id = 1
            AND team_id = $2
            AND status = 'active';
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
        INSERT INTO team_members (
            member_id,
            team_id,
            role_id,
            status
        )
        VALUES (
            $1,
            $2,
            $3,
            $4
        )
        RETURNING
            id,
            created_at
    `
	args := []any{teamMember.MemberId, teamMember.TeamId, teamMember.RoleId, teamMember.Status}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return tm.DB.QueryRowContext(ctx, query, args...).Scan(
		&teamMember.ID,
		&teamMember.CreatedAt,
	)
}

func (tm *teamMemberRepository) Update(memberId int64, member db.UpdateParams) error {
    query := `
        UPDATE
            team_members
        SET
            status = $1
        WHERE
            memberId = $2;
    `
	args := []any{
        member.Status,
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

func (tm *teamMemberRepository) UpdateRole(teamId int64, memberId int64, roleId int64) error {
	query := `
        UPDATE
            team_members 
        SET
            role_id = $1
        WHERE
            team_id = $2
            AND member_id = $3;
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
