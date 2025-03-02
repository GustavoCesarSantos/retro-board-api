package poll

type Handlers struct {
	CreatePoll *CreatePoll
	DeleteOption *DeleteOption
	ListAllPolls *ListAllPolls
	ListPoll *ListPoll
	ShowPollResult *ShowPollResult
	Vote *Vote
}

func NewHandlers(
	createPoll *CreatePoll,
	deleteOption *DeleteOption,
	listAllPolls *ListAllPolls,
	listPoll *ListPoll,
	showPollResult *ShowPollResult,
	vote *Vote,
) *Handlers {
	return &Handlers{
		CreatePoll: createPoll,
		DeleteOption: deleteOption,
		ListAllPolls: listAllPolls,
		ListPoll: listPoll,
		ShowPollResult: showPollResult,
		Vote: vote,
	}
}
