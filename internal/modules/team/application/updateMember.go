package application

import db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"

type IUpdateMember interface {
    Execute(memberId int64, status *string) error
}

type updateMember struct {
    repository db.ITeamMemberRepository
}

func NewUpdateMember(repository db.ITeamMemberRepository) IUpdateMember {
    return &updateMember{
        repository,
    }
}

func (um *updateMember) Execute(memberId int64, status *string) error {
    member := struct {
        Status *string
    }{
        Status: status,
    }
    return um.repository.Update(memberId, member)
}
