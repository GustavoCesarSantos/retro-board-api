package domain

import "time"

type Role struct {
    ID int64    
    Name string
}

type User struct {
    ID int64    
    Name string
    Email string
}

type TeamMember struct {
    ID int64
    TeamId int64
    User User
    Role Role
    Status string //active - deactivated - invited
    CreatedAt time.Time
    UpdatedAt *time.Time
}

func NewTeamMember(
    id int64,
    teamId int64,
    userId int64,
	roleId int64,
    status string,
) *TeamMember {
    return &TeamMember{
        ID: id,
        TeamId: teamId,
        User: User{ ID: userId },
        Role: Role{ ID: roleId },
        Status: status,
        CreatedAt: time.Now(),
    }
}
