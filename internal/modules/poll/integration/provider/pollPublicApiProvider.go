package provider

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
)

type pollPublicApiProvider struct {
    optionRepository db.IOptionRepository
    pollRepository db.IPollRepository
}

func NewPollPublicApiProvider(
    optionRepository db.IOptionRepository,
    pollRepository db.IPollRepository,
) interfaces.IPollApi {
    return &pollPublicApiProvider{
        optionRepository,
        pollRepository,
    }
}

func (ppa pollPublicApiProvider) FindAllOptionsByPollId(pollId int64) ([]*domain.Option, error) {
	return ppa.optionRepository.FindAllByPollId(pollId)
}

func (ppa pollPublicApiProvider) FindAllPollsByTeamId(teamId int64) ([]*domain.Poll, error) {
    return ppa.pollRepository.FindAllByTeamId(teamId)
}
