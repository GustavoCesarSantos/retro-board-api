package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type teamRepository struct {
	teams []domain.Team
}

func NewTeamRepository() db.ITeamRepository {
	return &teamRepository{
		teams: []domain.Team{
			*domain.NewTeam(1, "Time 1", 1),
			*domain.NewTeam(2, "Time 2", 1),
			*domain.NewTeam(3, "Time 3", 2),
		},
	}
}

func (tr *teamRepository) Delete(teamId int64, adminId int64,) error {
	i := 0
	for _, team := range tr.teams {
		if !(team.ID == teamId) {
			tr.teams[i] = team
			i++
		}
	}
	tr.teams = tr.teams[:i]
    return nil
}

func (tr *teamRepository) FindAllByAdminId(adminId int64) ([]*domain.Team, error) {
    var teams []*domain.Team
    for _, team := range tr.teams {
        if team.AdminId == adminId {
            teams = append(teams, &team)
        }
    }
    return teams, nil
}

func (tr *teamRepository) FindAllByMemberId(memberId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Team], error) {
    var teams []domain.Team
    for _, team := range tr.teams {
        if team.AdminId == memberId {
            teams = append(teams, team)
        }
    }
    return &utils.ResultPaginated[domain.Team]{
        Items: teams,
        NextCursor: 0,
    }, nil
}

func (tr *teamRepository) FindById(teamId int64, memberId int64) (*domain.Team, error) {
    for _, team := range tr.teams {
        if team.ID == teamId {
            return &team, nil
        }
    }
    return nil, utils.ErrRecordNotFound
}

func (tr *teamRepository) Save(team *domain.Team) error {
	tr.teams = append(tr.teams, *team)
    return nil
}
