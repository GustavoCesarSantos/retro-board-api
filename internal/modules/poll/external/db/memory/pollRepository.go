package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/interfaces"
)

type pollRepository struct {
	polls []domain.Poll
}

func NewPollRepository() db.IPollRepository {
	return &pollRepository{
		polls: []domain.Poll{
			*domain.NewPoll(1, 1, "Poll 1"),
			*domain.NewPoll(2, 1, "Poll 2",),
			*domain.NewPoll(3, 2, "Poll 1",),
		},
	}
}

func (pr *pollRepository) FindAllByTeamId(teamId int64) ([]*domain.Poll, error) {
    var polls []*domain.Poll
    for _, poll := range pr.polls {
        if poll.TeamId == teamId {
            polls = append(polls, &poll)
        }
    }
    return polls, nil
}

func (pr *pollRepository) FindById(pollId int64) (*domain.Poll, error) {
    for _, poll := range pr.polls {
        if poll.ID == pollId {
            return &poll, nil
        }
    }
    return nil, nil
}

func (pr *pollRepository) Save(poll *domain.Poll) error {
	pr.polls = append(pr.polls, *poll)
    return nil
}

