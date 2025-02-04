package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"
)

type IFindAllMembers interface {
    Execute(teamId int64) ([]*domain.TeamMember, error)
}

type findAllMembers struct {
    repository db.ITeamMemberRepository
}

func NewFindAllMembers(repository db.ITeamMemberRepository) IFindAllMembers {
    return &findAllMembers{
        repository,
    }
}

func (fam *findAllMembers) Execute(teamId int64) ([]*domain.TeamMember, error) {
    return fam.repository.FindAllByTeamId(teamId)
}
