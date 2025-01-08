package interfaces

import "net/http"

type IRoomManagerApi interface {
    AddUserToRoom(w http.ResponseWriter, r *http.Request, category string, roomId int64, userId int64) error
    BroadcastMessage(category string, roomId int64, message []byte)
    CloseConnection(category string, roomId int64, userId int64)
    ReadMessage(category string, roomId int64, userId int64) (int, []byte, error)
}
