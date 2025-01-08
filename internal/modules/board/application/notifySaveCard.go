package application

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type INotifySaveCard interface { 
    Execute(boardId int64, columnId int64, card *domain.Card) error
}

type notifySaveCard struct {
    roomManager interfaces.IRoomManagerApi
}

func NewNotifySaveCard(roomManager interfaces.IRoomManagerApi) INotifySaveCard {
    return &notifySaveCard{
        roomManager,
    }
}

func (nsc *notifySaveCard) Execute(boardId int64, columnId int64, card *domain.Card) error {
    cardJSON , err := json.Marshal(card)
    if err != nil {
        return errors.New("FAILED_TO_NOTIFY_CREATE_CARD_EVENT")
    }
    message := []byte(
        fmt.Sprintf(
            `{ "event": %s, "data": { "boardId": %d,  "columnId": %d, "card": %s  } }`,
            utils.CreateCardEvent,
            boardId,
            columnId,
            cardJSON,
        ),
    )
    nsc.roomManager.BroadcastMessage("boards", boardId, message) 
    return nil
}
