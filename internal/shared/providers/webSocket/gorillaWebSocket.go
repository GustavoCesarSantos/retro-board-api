package providers

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
	"github.com/gorilla/websocket"
)

type RoomManager struct {
	rooms map[string]map[int64]map[int64]*websocket.Conn
    mu    sync.Mutex
}

type gorillaWebSocket struct {
    roomManager RoomManager
}

func NewGorillaWebSocket() interfaces.IRoomManagerApi {
    return &gorillaWebSocket{
        roomManager: RoomManager{
            rooms: map[string]map[int64]map[int64]*websocket.Conn{
                "boards": make(map[int64]map[int64]*websocket.Conn),
                "polls": make(map[int64]map[int64]*websocket.Conn),
            },
        },
    }
}

func createConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
    var upgrader = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true
        },
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
    }
    conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
        slog.Error(err.Error(), "meta", fmt.Sprintf("%s", utils.Envelope{
            "file": "gorillaWebSocket.go",
            "func": "gorillaWebSocket.createConnection",
            "line": 42,
	    }))
		return nil, err
	}
    return conn, nil
}

func (gws *gorillaWebSocket) AddUserToRoom(w http.ResponseWriter, r *http.Request, category string, roomId int64, userId int64) error {
    conn, err := createConnection(w, r)
    if err != nil {
        slog.Error(err.Error(), "meta", fmt.Sprintf("%s", utils.Envelope{
            "file": "gorillaWebSocket.go",
            "func": "gorillaWebSocket.AddUserToRoom",
            "line": 55,
	    }))
        return err
    }
    gws.roomManager.mu.Lock()
	if _, ok := gws.roomManager.rooms[category]; !ok {
		gws.roomManager.rooms[category] = make(map[int64]map[int64]*websocket.Conn)
	}
    if _, ok := gws.roomManager.rooms[category][roomId]; !ok {
		gws.roomManager.rooms[category][roomId] = make(map[int64]*websocket.Conn)
	}
	gws.roomManager.rooms[category][roomId][userId] = conn
	gws.roomManager.mu.Unlock()
    slog.Info(fmt.Sprintf("Conexão WebSocket estabelecida para a sala: %d", roomId))
    return nil
}

func (gws *gorillaWebSocket) BroadcastMessage(category string, roomId int64, message []byte) {
	gws.roomManager.mu.Lock()
	defer gws.roomManager.mu.Unlock()
	for userId := range gws.roomManager.rooms[category][roomId] {
        conn := gws.roomManager.rooms[category][roomId][userId]
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
            slog.Error(err.Error(), "meta", fmt.Sprintf("%s", utils.Envelope{
                "file": "gorillaWebSocket.go",
                "func": "gorillaWebSocket.BroadcastMessage",
                "line": 82,
	        }))
			conn.Close()
			delete(gws.roomManager.rooms[category][roomId], userId)
		}
	}
}

func (gws *gorillaWebSocket) CloseConnection(category string, roomId int64, userId int64) {
    gws.roomManager.mu.Lock()
    conn := gws.roomManager.rooms[category][roomId][userId]
    delete(gws.roomManager.rooms[category][roomId], userId)
    if len(gws.roomManager.rooms[category][roomId]) == 0 {
        delete(gws.roomManager.rooms[category], roomId)
    }
    if len(gws.roomManager.rooms[category]) == 0 {
        delete(gws.roomManager.rooms, category)
    }
    gws.roomManager.mu.Unlock()
    conn.Close()
    slog.Info(fmt.Sprintf("Conexão WebSocket encerrada com a sala: %d", roomId))
}

func (gws *gorillaWebSocket) ReadMessage(category string, roomId int64, userId int64) (int, []byte, error) {
    conn := gws.roomManager.rooms[category][roomId][userId]
    return conn.ReadMessage()
}

