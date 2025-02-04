package dtos

import "time"

type ListAllMembersResponse struct {
	ID int64 `json:"id" example:"1"`
	TeamId int64 `json:"team_id" example:"1"`
	MemberId int64 `json:"member_id" example:"1"`
	RoleId int64 `json:"role_id" example:"1"`
	Status string `json:"status" example:"team1"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}

func NewListAllMembersResponse(
	id int64, 
    teamId int64,
    memberId int64,
    roleId int64,
	status string, 
	createdAt time.Time,
	updatedAt *time.Time,
) *ListAllMembersResponse {
	return &ListAllMembersResponse {
		ID: id,
        TeamId: teamId,
        MemberId: memberId,
        RoleId: roleId,
        Status: status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
