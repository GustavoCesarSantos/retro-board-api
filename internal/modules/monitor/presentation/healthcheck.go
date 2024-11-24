package monitor

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/configs"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type healthcheck struct {}

func NewHealthcheck() *healthcheck {
    return &healthcheck{}
}

func (hc *healthcheck) Handle(w http.ResponseWriter, r *http.Request) {
	serverConfig := configs.LoadServerConfig()
	data := utils.Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": serverConfig.Env,
		},
	}
	err := utils.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		utils.ServerErrorResponse(w, r, err)
	}
}
