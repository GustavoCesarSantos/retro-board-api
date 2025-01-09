package monitor

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/configs"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type healthcheck struct {}

func NewHealthcheck() *healthcheck {
    return &healthcheck{}
}

type HealthCheckEnvelope struct {
	HealthCheck dtos.HealthCheckResponse `json:"health_check"`
}

// Handle performs a health check for the application.
// @Summary Application health check
// @Description Returns the current health status of the application and environment information.
// @Tags Monitor
// @Produce json
// @Success 200 {object} monitor.HealthCheckEnvelope "Health check result"
// @Failure 500 {object} utils.ErrorEnvelope "Internal server error"
// @Router /healthcheck [get]
func (hc *healthcheck) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "healthcheck.go",
		"func": "healthcheck.Handle",
		"line": 0,
	}
	serverConfig := configs.LoadServerConfig()
	response := dtos.NewHealthCheckResponse("available", serverConfig.Env)
	err := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"health_check": response}, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err, metadataErr)
	}
}
