package realtime

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type connectToBoardRoom struct {
    roomManager interfaces.IRoomManagerApi
}

func NewConnectToBoardRoom (
    roomManager interfaces.IRoomManagerApi,
) *connectToBoardRoom {
    return &connectToBoardRoom{
        roomManager,
    }
}

func(cbr *connectToBoardRoom) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "connectToBoardRoom.go",
		"func": "connectToBoardRoom.Handle",
		"line": 0,
	}
	user := utils.ContextGetUser(r)
	boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
	if boardIdErr != nil {
        utils.BadRequestResponse(w, r, boardIdErr, metadataErr)
		return
	}
    addConnectionErr := cbr.roomManager.AddUserToRoom(w, r, "boards", boardId, user.ID)
    if addConnectionErr != nil {
        utils.BadRequestResponse(w, r, addConnectionErr, metadataErr)
        return
    }
    for {
		_, _, err := cbr.roomManager.ReadMessage("boards", boardId, user.ID)
		if err != nil {
            cbr.roomManager.CloseConnection("boards", boardId, user.ID)
			break
		}
	}
}
