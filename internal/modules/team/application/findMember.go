package application

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"
)

type IFindMember interface {
    Execute(memberId int64) (*domain.TeamMember, error)
}

type findMember struct {
    repository db.ITeamMemberRepository
}

func NewFindMember(repository db.ITeamMemberRepository) IFindMember {
    return &findMember{
        repository,
    }
}

func (fm *findMember) Execute(memberId int64) (*domain.TeamMember, error) {
    return fm.repository.FindById(memberId)
}
