package team

type Handlers struct {
	AddMemberToTeam *AddMemberToTeam
	ChangeMemberRole *ChangeMemberRole
	CreateTeam *CreateTeam
	DeleteTeam *DeleteTeam
	EditMember *EditMember
	EditTeam *EditTeam
	ListAllMembers *ListAllMembers
	ListAllTeams *ListAllTeams
	RemoveMemberFromTeam *RemoveMemberFromTeam
	ShowTeam *ShowTeam
}

func NewHandlers (
	addMemberToTeam *AddMemberToTeam,
	changeMemberRole *ChangeMemberRole,
	createTeam *CreateTeam,
	deleteTeam *DeleteTeam,
	editMember *EditMember,
	editTeam *EditTeam,
	listAllMembers *ListAllMembers,
	listAllTeams *ListAllTeams,
	removeMemberFromTeam *RemoveMemberFromTeam,
	showTeam *ShowTeam,
) *Handlers {
	return &Handlers{
		AddMemberToTeam: addMemberToTeam,
		ChangeMemberRole: changeMemberRole,
		CreateTeam: createTeam,
		DeleteTeam: deleteTeam,
		EditMember: editMember,
		EditTeam: editTeam,
		ListAllMembers: listAllMembers,
		ListAllTeams: listAllTeams,
		RemoveMemberFromTeam: removeMemberFromTeam,
		ShowTeam: showTeam,
	}
}
