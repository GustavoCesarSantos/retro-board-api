package dtos

type SystemInfo struct {
	Environment string `json:"environment" example:"develop"`
}

type AddMemberToTeamRequest struct {
    Email string `json:"email" example:"useremail@foo.bar"`
    RoleId int64 `json:"role_id" example:"2"`
}
