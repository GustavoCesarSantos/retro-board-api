package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/domain"

type ITeamMemberRepository interface {
	Delete(teamId int64, memberId int64)
	Save(team domain.TeamMember)
	UpdateRole(teamId int64, memberId int64, role int64)
}

type teamMemberRepository struct {
	teamMembers []domain.TeamMember
}

func NewTeamMemberRepository() ITeamMemberRepository {
	return &teamMemberRepository{
		teamMembers: []domain.TeamMember{
			*domain.NewTeamMember(1, 1, 1, 1),
			*domain.NewTeamMember(2, 1, 2, 2),
			*domain.NewTeamMember(3, 1, 3, 2),
			*domain.NewTeamMember(4, 1, 4, 2),
		},
	}
}

func (tm *teamMemberRepository) Delete(teamId int64, memberId int64,) {
	i := 0
	for _, member := range tm.teamMembers {
		if !(member.TeamId == teamId && member.MemberId == memberId) {
			tm.teamMembers[i] = member
			i++
		}
	}
	tm.teamMembers = tm.teamMembers[:i]
}

func (tm *teamMemberRepository) Save(teamMember domain.TeamMember) {
	tm.teamMembers = append(tm.teamMembers, teamMember)
}

func (tm *teamMemberRepository) UpdateRole(teamId int64, memberId int64, role int64) {
	for i := range tm.teamMembers {
		if tm.teamMembers[i].TeamId == teamId && tm.teamMembers[i].MemberId == memberId {
			tm.teamMembers[i].Role = role
			break
		}
	}
}