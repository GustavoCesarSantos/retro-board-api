package providers

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/gorilla/websocket"
)

type BoardManager struct {
	rooms map[int64]map[int64]*websocket.Conn
    mu    sync.Mutex
}

type gorillaWebSocket struct {
    boardManager BoardManager
}

func NewGorillaWebSocket() interfaces.IRoomManagerIntegration {
    return &gorillaWebSocket{
        boardManager: BoardManager{
            rooms: make(map[int64]map[int64]*websocket.Conn),
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

func (gws *gorillaWebSocket) AddUserToRoom(w http.ResponseWriter, r *http.Request, boardId int64, userId int64) error {
    conn, err := createConnection(w, r)
    if err != nil {
        return err
    }
    gws.boardManager.mu.Lock()
	if _, ok := gws.boardManager.rooms[boardId]; !ok {
		gws.boardManager.rooms[boardId] = make(map[int64]*websocket.Conn)
	}
	gws.boardManager.rooms[boardId][userId] = conn
	gws.boardManager.mu.Unlock()
    slog.Info("Conexão WebSocket estabelecida para o board: %s", boardId)
    return nil
}

func (gws *gorillaWebSocket) BroadcastToBoard(boardId int64, message []byte) {
	gws.boardManager.mu.Lock()
	defer gws.boardManager.mu.Unlock()
	for userId := range gws.boardManager.rooms[boardId] {
        conn := gws.boardManager.rooms[boardId][userId]
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			slog.Info("Erro ao enviar mensagem para WebSocket: %v", err)
			conn.Close()
			delete(gws.boardManager.rooms[boardId], userId)
		}
	}
}

func (gws *gorillaWebSocket) CloseConnection(boardId int64, userId int64) {
    gws.boardManager.mu.Lock()
    conn := gws.boardManager.rooms[boardId][userId]
    delete(gws.boardManager.rooms[boardId], userId)
    if len(gws.boardManager.rooms[boardId]) == 0 {
        delete(gws.boardManager.rooms, boardId)
    }
    gws.boardManager.mu.Unlock()
    conn.Close()
    slog.Info("Conexão WebSocket encerrada para o board: %s", boardId)
}

func (gws *gorillaWebSocket) ReadMessage(boardId int64, userId int64) (int, []byte, error) {
    conn := gws.boardManager.rooms[boardId][userId]
    return conn.ReadMessage()
}

