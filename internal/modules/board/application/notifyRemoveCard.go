package application

import (
	"fmt"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type INotifyRemoveCard interface { 
    Execute(boardId int64, columnId int64, cardId int64)
}

type notifyRemoveCard struct {
    roomManager interfaces.IRoomManagerApi
}

func NewNotifyRemoveCard(roomManager interfaces.IRoomManagerApi) INotifyRemoveCard {
    return &notifyRemoveCard{
        roomManager,
    }
}

func (nmc *notifyRemoveCard) Execute(boardId int64, columnId int64, cardId int64) {
    message := []byte(
        fmt.Sprintf(
            `{ "event": %s, "data": { "boardId": %d,  "columnId": %d, "cardId": %d  } }`,
            utils.RemoveCardEvent,
            boardId,
            columnId,
            cardId,
        ),
    )
    nmc.roomManager.BroadcastMessage("boards", boardId, message) 
}
