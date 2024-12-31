package interfaces

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"

type ITeamMemberApi interface {
	FindAllByTeamId(teamId int64) ([]*domain.TeamMember, error)
}