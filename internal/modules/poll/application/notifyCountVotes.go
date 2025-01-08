package application

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type INotifyCountVotes interface { 
    Execute(pollId int64, result *CountVotesResult) error
}

type notifyCountVotes struct {
    roomManager interfaces.IRoomManagerApi
}

func NewNotifyCountVotes(roomManager interfaces.IRoomManagerApi) INotifyCountVotes {
    return &notifyCountVotes{
        roomManager,
    }
}

func (nsv *notifyCountVotes) Execute(pollId int64, result *CountVotesResult) error {
    resultJSON , err := json.Marshal(result)
    if err != nil {
        return errors.New("FAILED_TO_NOTIFY_POLL_RESULT_EVENT")
    }
    message := []byte(
        fmt.Sprintf(
            `{ "event": %s, "data": { "pollId": %d,  "result": %s } }`,
            utils.VoteEvent,
            pollId,
            resultJSON,
        ),
    )
    nsv.roomManager.BroadcastMessage("polls", pollId, message) 
    return nil
}
