package providers

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/gorilla/websocket"
)

type BoardManager struct {
	rooms map[string]map[int64]map[int64]*websocket.Conn
    mu    sync.Mutex
}

type gorillaWebSocket struct {
    boardManager BoardManager
}

func NewGorillaWebSocket() interfaces.IRoomManagerApi {
    return &gorillaWebSocket{
        boardManager: BoardManager{
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
    }
    conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
    return conn, nil
}

func (gws *gorillaWebSocket) AddUserToRoom(w http.ResponseWriter, r *http.Request, category string, roomId int64, userId int64) error {
    conn, err := createConnection(w, r)
    if err != nil {
        return err
    }
    gws.boardManager.mu.Lock()
	if _, ok := gws.boardManager.rooms[category]; !ok {
		gws.boardManager.rooms[category] = make(map[int64]map[int64]*websocket.Conn)
	}
    if _, ok := gws.boardManager.rooms[category][roomId]; !ok {
		gws.boardManager.rooms[category][roomId] = make(map[int64]*websocket.Conn)
	}
	gws.boardManager.rooms[category][roomId][userId] = conn
	gws.boardManager.mu.Unlock()
    slog.Info(fmt.Sprintf("Conexão WebSocket estabelecida para a sala: %d", roomId))
    return nil
}

func (gws *gorillaWebSocket) BroadcastMessage(category string, roomId int64, message []byte) {
	gws.boardManager.mu.Lock()
	defer gws.boardManager.mu.Unlock()
	for userId := range gws.boardManager.rooms[category][roomId] {
        conn := gws.boardManager.rooms[category][roomId][userId]
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
            slog.Info(fmt.Sprintf("Erro ao enviar mensagem para WebSocket: %v", err))
			conn.Close()
			delete(gws.boardManager.rooms[category][roomId], userId)
		}
	}
}

func (gws *gorillaWebSocket) CloseConnection(category string, roomId int64, userId int64) {
    gws.boardManager.mu.Lock()
    conn := gws.boardManager.rooms[category][roomId][userId]
    delete(gws.boardManager.rooms[category][roomId], userId)
    if len(gws.boardManager.rooms[category][roomId]) == 0 {
        delete(gws.boardManager.rooms[category], roomId)
    }
    if len(gws.boardManager.rooms[category]) == 0 {
        delete(gws.boardManager.rooms, category)
    }
    gws.boardManager.mu.Unlock()
    conn.Close()
    slog.Info(fmt.Sprintf("Conexão WebSocket encerrada com a sala: %d", roomId))
}

func (gws *gorillaWebSocket) ReadMessage(category string, roomId int64, userId int64) (int, []byte, error) {
    conn := gws.boardManager.rooms[category][roomId][userId]
    return conn.ReadMessage()
}

