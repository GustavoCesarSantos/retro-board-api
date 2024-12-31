package provider

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
)

type teamMemberPublicApiProvider struct {
    repository db.ITeamMemberRepository
}

func NewTeamMemberPublicApiProvider(repository db.ITeamMemberRepository) interfaces.ITeamMemberApi {
    return &teamMemberPublicApiProvider{
        repository,
    }
}

func (tmpa teamMemberPublicApiProvider) FindAllByTeamId(teamId int64) ([]*domain.TeamMember, error) {
    return tmpa.repository.FindAllByTeamId(teamId)
}
