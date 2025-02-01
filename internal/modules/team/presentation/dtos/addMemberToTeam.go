package dtos

type AddMemberToTeamRequest struct {
    Email string `json:"email" example:"useremail@foo.bar"`
    RoleId int64 `json:"role_id" example:"2"`
}
