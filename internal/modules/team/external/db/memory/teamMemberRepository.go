package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type teamMemberRepository struct {
	teamMembers []domain.TeamMember
}

func NewTeamMemberRepository() db.ITeamMemberRepository {
	return &teamMemberRepository{
		teamMembers: []domain.TeamMember{
			*domain.NewTeamMember(1, 1, 1, 1, "active"),
			*domain.NewTeamMember(2, 1, 2, 2, "active"),
			*domain.NewTeamMember(3, 1, 3, 2, "active"),
			*domain.NewTeamMember(4, 1, 4, 2, "active"),
		},
	}
}

func (tm *teamMemberRepository) Delete(teamId int64, memberId int64,) error {
	i := 0
	for _, member := range tm.teamMembers {
		if !(member.TeamId == teamId && member.User.ID == memberId) {
			tm.teamMembers[i] = member
			i++
		}
	}
	tm.teamMembers = tm.teamMembers[:i]
    return nil
}

func (tm *teamMemberRepository) FindAllByTeamId(teamId int64) ([]*domain.TeamMember, error) {
	var teamMembers []*domain.TeamMember
    for _, teamMember := range tm.teamMembers {
        if teamMember.TeamId == teamId && teamMember.Status == "active" {
            teamMembers = append(teamMembers, &teamMember)
        }
    }
    return teamMembers, nil
}

func (tm *teamMemberRepository) FindById(memberId int64) (*domain.TeamMember, error) {
    for _, teamMember := range tm.teamMembers {
        if teamMember.ID == memberId {
            return &teamMember, nil
        }
    }
    return nil, utils.ErrRecordNotFound
}

func (tm *teamMemberRepository) FindTeamAdminByMemberId(teamId int64, memberId int64) (*domain.TeamMember, error) {
    for _, teamMember := range tm.teamMembers {
        if teamMember.TeamId == teamId && teamMember.User.ID == memberId && teamMember.Role.ID == 1 {
            return &teamMember, nil
        }
    }
    return nil, utils.ErrRecordNotFound
}

func (tm *teamMemberRepository) Save(teamMember *domain.TeamMember) error {
	tm.teamMembers = append(tm.teamMembers, *teamMember)
    return nil
}

func (tm *teamMemberRepository) Update(memberId int64, member db.UpdateMemberParams) error {
	for i := range tm.teamMembers {
		if tm.teamMembers[i].User.ID == memberId {
            if member.Status != nil {
                tm.teamMembers[i].Status = *member.Status
            }
			break
		}
	}
    return nil
}

func (tm *teamMemberRepository) UpdateRole(teamId int64, memberId int64, roleId int64) error {
	for i := range tm.teamMembers {
		if tm.teamMembers[i].TeamId == teamId && tm.teamMembers[i].User.ID == memberId {
			tm.teamMembers[i].Role.ID = roleId
			break
		}
	}
    return nil
}
