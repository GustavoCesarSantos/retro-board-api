package interfaces

import "net/http"

type IRoomManagerIntegration interface {
    AddUserToRoom(w http.ResponseWriter, r *http.Request, boardId int64, userId int64) error
    BroadcastToBoard(boardId int64, message []byte)
    CloseConnection(boardId int64, userId int64)
    ReadMessage(boardId int64, userId int64) (int, []byte, error)
}
