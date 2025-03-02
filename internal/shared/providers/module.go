package providers

import (
	"go.uber.org/fx"

	websocket "github.com/GustavoCesarSantos/retro-board-api/internal/shared/providers/webSocket"
)

var Module = fx.Module(
	"providers",
	fx.Provide(
		websocket.NewGorillaWebSocket,
	),
)