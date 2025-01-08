package application

import (
	"fmt"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type INotifySaveVote interface { 
    Execute(pollId int64, optionId int64)
}

type notifySaveVote struct {
    roomManager interfaces.IRoomManagerApi
}

func NewNotifySaveVote(roomManager interfaces.IRoomManagerApi) INotifySaveVote {
    return &notifySaveVote{
        roomManager,
    }
}

func (nsv *notifySaveVote) Execute(pollId int64, optionId int64) {
    message := []byte(
        fmt.Sprintf(
            `{ "event": %s, "data": { "pollId": %d,  "optionId": %d } }`,
            utils.VoteEvent,
            pollId,
            optionId,
        ),
    )
    nsv.roomManager.BroadcastMessage("polls", pollId, message) 
}
