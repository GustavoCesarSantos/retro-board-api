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

type teamRepository struct {
    DB *sql.DB
}

func NewTeamRepository(db *sql.DB) db.ITeamRepository {
	return &teamRepository{
        DB: db,
	}
}

func (tr *teamRepository) Delete(teamId int64, adminId int64) error {
    query := `
        DELETE FROM 
            teams t
        USING 
            team_members tm
        WHERE 
            t.id = $1 
            AND tm.team_id = t.id 
            AND tm.member_id = $2
            AND tm.role_id = 1;
    `
	args := []any{teamId, adminId}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := tr.DB.ExecContext(ctx, query, args...)
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

func (tr *teamRepository) FindAllByAdminId(adminId int64) ([]*domain.Team, error) {
    query := `
        SELECT
            t.id,
            t.name,
            t.created_at,
            t.updated_at
        FROM
            teams t
        INNER JOIN
            team_members tm
        ON
            tm.team_id = t.id
        WHERE
            tm.member_id = $1
            AND tm.role_id = 1;
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := tr.DB.QueryContext(ctx, query, adminId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    teams := []*domain.Team{}
    for rows.Next() {
        var team domain.Team
        err := rows.Scan(
            &team.ID,
            &team.Name,
            &team.CreatedAt,
            &team.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        teams = append(teams, &team)
    }
    if rowsErr := rows.Err(); rowsErr != nil {
        return nil, rowsErr
    }
    return teams, nil
}

func (tr *teamRepository) FindAllByMemberId(memberId int64) ([]*domain.Team, error) {
    query := `
        SELECT
            t.id,
            t.name,
            t.created_at,
            t.updated_at
        FROM
            teams t
        INNER JOIN
            team_members tm
        ON
            tm.team_id = t.id
        WHERE
            tm.member_id = $1;
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := tr.DB.QueryContext(ctx, query, memberId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    teams := []*domain.Team{}
    for rows.Next() {
        var team domain.Team
        err := rows.Scan(
            &team.ID,
            &team.Name,
            &team.CreatedAt,
            &team.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        teams = append(teams, &team)
    }
    if rowsErr := rows.Err(); rowsErr != nil {
        return nil, rowsErr
    }
    return teams, nil
}

func (tr *teamRepository) FindById(teamId int64, memberId int64) (*domain.Team, error) {
    query := `
        SELECT
            t.id,
            t.name,
            t.created_at,
            t.updated_at
        FROM
            teams t
        INNER JOIN
            team_members tm
        ON
            tm.team_id = t.id
        WHERE
            t.id = $1
            AND tm.member_id = $2;
    `
	args := []any{teamId, memberId}
	var team domain.Team
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := tr.DB.QueryRowContext(ctx, query, args...).Scan(
		&team.ID,
        &team.Name,
		&team.CreatedAt,
		&team.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, utils.ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &team, nil
}

func (tr *teamRepository) Save(team *domain.Team) error {
    query := `
        INSERT INTO teams (
            name
        )
        VALUES (
            $1
        )
        RETURNING
            id,
            created_at
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
    return tr.DB.QueryRowContext(ctx, query, team.Name).Scan(
        &team.ID,
        &team.CreatedAt,
    )
}
