package application

import (
	"fmt"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type INotifyUpdateCard interface { 
    Execute(boardId int64, columnId int64, cardText string)
}

type notifyUpdateCard struct {
    roomManager interfaces.IRoomManagerApi
}

func NewNotifyUpdateCard(roomManager interfaces.IRoomManagerApi) INotifyUpdateCard {
    return &notifyUpdateCard{
        roomManager,
    }
}

func (nsc *notifyUpdateCard) Execute(boardId int64, columnId int64, cardText string) {
    message := []byte(
        fmt.Sprintf(
            `{ "event": %s, "data": { "boardId": %d,  "columnId": %d, "cardText": %s  } }`,
            utils.EditCardEvent,
            boardId,
            columnId,
            cardText,
        ),
    )
    nsc.roomManager.BroadcastMessage("boards", boardId, message) 
}
