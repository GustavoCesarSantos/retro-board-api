package dtos

import "time"

type Role struct {
    ID int64 `json:"role_id" example:"1"`
    Name string `json:"role_name" example:"admin"`
}

type User struct {
    Name string `json:"name" example:"nome do usuário"`
    Email string `json:"email" example:"email do usuário"`
}

type ListAllMembersResponse struct {
	ID int64 `json:"id" example:"1"`
	TeamId int64 `json:"team_id" example:"1"`
    User User `json:"user"`
	Role Role `json:"role"`
	Status string `json:"status" example:"team1"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

func NewListAllMembersResponse(
	id int64, 
    teamId int64,
    userName string,
    userEmail string,
    roleId int64,
    roleName string,
	status string, 
	createdAt time.Time,
	updatedAt *time.Time,
) *ListAllMembersResponse {
	return &ListAllMembersResponse {
		ID: id,
        TeamId: teamId,
        User: User{ Name: userName, Email: userEmail },
        Role: Role{ ID: roleId, Name: roleName },
        Status: status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
