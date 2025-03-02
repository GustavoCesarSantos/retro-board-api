package team

import (
	"go.uber.org/fx"

	teamApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	teamDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/nativeSql"
	teamMemberProvider "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/integration/provider"
	team "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation/handlers"
)

var Module = fx.Module(
	"team",
	fx.Provide(
		// Repositories
		teamDb.NewTeamRepository,
		teamDb.NewTeamMemberRepository,

		// Providers
		teamMemberProvider.NewTeamMemberPublicApiProvider,
	
		// Applications
		teamApplication.NewEnsureAdminMembership,
		teamApplication.NewFindAllMembers,
		teamApplication.NewFindAllTeams,
		teamApplication.NewFindTeam,
		teamApplication.NewRemoveMember,
		teamApplication.NewRemoveTeam,
		teamApplication.NewSaveMember,
		teamApplication.NewSaveTeam,
		teamApplication.NewUpdateMember,
		teamApplication.NewUpdateRole,
		teamApplication.NewUpdateTeam,
		teamApplication.NewFindMemberInfoByEmail,	

		// Handlers
		team.NewAddMemberToTeam,
		team.NewChangeMemberRole,
		team.NewCreateTeam,
		team.NewDeleteTeam,
		team.NewEditMember,
		team.NewEditTeam,
		team.NewListAllMembers,
		team.NewListAllTeams,
		team.NewRemoveMemberFromTeam,
		team.NewShowTeam,

		// Handlers
		team.NewHandlers,
	),
)