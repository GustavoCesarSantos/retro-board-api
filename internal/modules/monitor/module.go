package monitor

import (
	"go.uber.org/fx"

	monitor "github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor/presentation/handlers"
)

var Module = fx.Module(
	"monitor",
	fx.Provide(
		monitor.NewHealthcheck,

		// Handlers
		monitor.NewHandlers,
	),
)