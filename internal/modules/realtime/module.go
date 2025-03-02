package realtime

import (
	"go.uber.org/fx"

	realtime "github.com/GustavoCesarSantos/retro-board-api/internal/modules/realtime/presentation/handlers"
)

var Module = fx.Module(
	"realtime",
	fx.Provide(
		realtime.NewConnectToBoardRoom,

		// Handlers
		realtime.NewHandlers,
	),
)