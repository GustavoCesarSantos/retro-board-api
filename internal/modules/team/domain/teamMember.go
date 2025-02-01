package domain

import "time"

type TeamMember struct {
    ID int64
    TeamId int64
    MemberId int64
	RoleId int64
    Status string
    CreatedAt time.Time
    UpdatedAt *time.Time
}

func NewTeamMember(
    id int64,
    teamId int64,
    memberId int64,
	roleId int64,
    status string,
) *TeamMember {
    return &TeamMember{
        ID: id,
        TeamId: teamId,
		MemberId: memberId,
		RoleId: roleId,
        Status: status,
        CreatedAt: time.Now(),
    }
}
