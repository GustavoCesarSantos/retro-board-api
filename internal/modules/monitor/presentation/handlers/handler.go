package monitor

type Handlers struct {
	Healthcheck *Healthcheck
}

func NewHandlers(
	healthcheck *Healthcheck,
) *Handlers {
	return &Handlers{
		Healthcheck: healthcheck,
	}
}
