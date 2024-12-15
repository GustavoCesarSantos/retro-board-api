package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"

type ITeamRepository interface {
    Delete(teamId int64, adminId int64) error
    FindAllByAdminId(adminId int64) ([]*domain.Team, error)
    FindAllByMemberId(memberId int64) ([]*domain.Team, error)
    FindById(teamId int64, memberId int64) (*domain.Team, error)
	Save(team *domain.Team) error
}
