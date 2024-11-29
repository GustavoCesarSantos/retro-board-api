package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"

type ITeamRepository interface {
    FindAllByAdminId(adminId int64) []*domain.Team
    FindById(teamId int64, adminId int64) *domain.Team
	Save(team domain.Team)
}

type teamRepository struct {
	teams []domain.Team
}

func NewTeamRepository() ITeamRepository {
	return &teamRepository{
		teams: []domain.Team{
			*domain.NewTeam(1, "Time 1", 1),
			*domain.NewTeam(2, "Time 2", 1),
			*domain.NewTeam(3, "Time 3", 2),
		},
	}
}

func (tr *teamRepository) FindAllByAdminId(adminId int64) []*domain.Team {
    var teams []*domain.Team
    for _, team := range tr.teams {
        if team.AdminId == adminId {
            teams = append(teams, &team)
        }
    }
    return teams
}

func (tr *teamRepository) FindById(teamId int64, adminId int64) *domain.Team {
    for _, team := range tr.teams {
        if team.ID == teamId && team.AdminId == adminId {
            return &team
        }
    }
    return nil
}

func (tr *teamRepository) Save(team domain.Team) {
	tr.teams = append(tr.teams, team)
}
