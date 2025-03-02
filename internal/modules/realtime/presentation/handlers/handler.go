package realtime

type Handlers struct {
	ConnectToBoardRoom *ConnectToBoardRoom
}

func NewHandlers(
	connectToBoardRoom *ConnectToBoardRoom,
) *Handlers {
	return &Handlers{
		ConnectToBoardRoom: connectToBoardRoom,
	}
}
