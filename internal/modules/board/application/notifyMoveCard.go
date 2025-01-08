package application

import (
	"fmt"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type INotifyMoveCard interface { 
    Execute(boardId int64, columnId int64, cardId int64)
}

type notifyMoveCard struct {
    roomManager interfaces.IRoomManagerApi
}

func NewNotifyMoveCard(roomManager interfaces.IRoomManagerApi) INotifyMoveCard {
    return &notifyMoveCard{
        roomManager,
    }
}

func (nmc *notifyMoveCard) Execute(boardId int64, columnId int64, cardId int64) {
    message := []byte(
        fmt.Sprintf(
            `{ "event": %s, "data": { "boardId": %d,  "columnId": %d, "card": %d  } }`,
            utils.MoveCardEvent,
            boardId,
            columnId,
            cardId,
        ),
    )
    nmc.roomManager.BroadcastMessage("boards", boardId, message) 
}
