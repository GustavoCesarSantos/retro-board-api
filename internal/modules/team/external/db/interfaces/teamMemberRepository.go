package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"

type ITeamMemberRepository interface {
	Delete(teamId int64, memberId int64) error
	FindAllByTeamId(teamId int64) ([]*domain.TeamMember, error)
	FindTeamAdminByMemberId(teamId int64, memberId int64) (*domain.TeamMember, error)
	Save(teamMember *domain.TeamMember) error
	UpdateRole(teamId int64, memberId int64, roleId int64) error
}
