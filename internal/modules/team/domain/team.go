package domain

import "time"

type Team struct {
    ID int64
    Name string
    AdminId int64
    CreatedAt time.Time
    UpdatedAt time.Time
}

func NewTeam(
    id int64,
    name string,
    adminId int64,
) *Team {
    return &Team{
        ID: id,
        Name: name,
        AdminId: adminId,
        CreatedAt: time.Now(),
    }
}
