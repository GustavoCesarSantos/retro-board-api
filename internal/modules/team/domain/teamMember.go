package domain

import "time"

type TeamMember struct {
    ID int64
    TeamId int64
    MemberId int64
	Role int64
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewTeamMember(
    id int64,
    teamId int64,
    memberId int64,
	role int64,
) *TeamMember {
    return &TeamMember{
        ID: id,
        TeamId: teamId,
		MemberId: memberId,
		Role: role,
        CreatedAt: time.Now(),
    }
}
