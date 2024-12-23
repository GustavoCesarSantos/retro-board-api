package dtos

type SystemInfo struct {
	Environment string `json:"environment" example:"develop"`
}

type AddMemberToTeamRequest struct {
	MemberId int64 `json:"memberId" example:"1"`
    RoleId int64 `json:"roleId" example:"2"`
}