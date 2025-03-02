package poll

import (
	"go.uber.org/fx"

	pollApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/application"
	pollDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/nativeSql"
	pollProvider "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/integration/provider"
	poll "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/presentation/handlers"
)

var Module = fx.Module(
	"poll",
	fx.Provide(
		// Repositories
		pollDb.NewOptionRepository,
		pollDb.NewPollRepository,
		pollDb.NewVoteRepository,

		// Providers
		pollProvider.NewPollPublicApiProvider,
		
		// Applications
		pollApplication.NewCountVotesByPollId,
		pollApplication.NewFindAllPolls,
		pollApplication.NewFindPoll,
		pollApplication.NewNotifyCountVotes,
		pollApplication.NewNotifySaveVote,
		pollApplication.NewRemoveOption,
		pollApplication.NewSaveOption,
		pollApplication.NewSavePoll,
		pollApplication.NewSaveVote,

		// Handlers
		poll.NewCreatePoll,
		poll.NewDeleteOption,
		poll.NewListAllPolls,
		poll.NewListPoll,
		poll.NewShowPollResult,
		poll.NewVote,

		// Handlers
		poll.NewHandlers,
	),
)