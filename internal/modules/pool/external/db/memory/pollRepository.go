package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/domain"

type IPollRepository interface {
    FindAllByTeamId(teamId int64) []*domain.Poll
    FindById(pollId int64) *domain.Poll
	Save(poll domain.Poll)
}

type pollRepository struct {
	polls []domain.Poll
}

func NewPollRepository() IPollRepository {
	return &pollRepository{
		polls: []domain.Poll{
			*domain.NewPoll(1, 1, "Poll 1"),
			*domain.NewPoll(2, 1, "Poll 2",),
			*domain.NewPoll(3, 2, "Poll 1",),
		},
	}
}

func (pr *pollRepository) FindAllByTeamId(teamId int64) []*domain.Poll {
    var polls []*domain.Poll
    for _, poll := range pr.polls {
        if poll.TeamId == teamId {
            polls = append(polls, &poll)
        }
    }
    return polls
}

func (pr *pollRepository) FindById(pollId int64) *domain.Poll {
    for _, poll := range pr.polls {
        if poll.ID == pollId {
            return &poll
        }
    }
    return nil
}

func (pr *pollRepository) Save(poll domain.Poll) {
	pr.polls = append(pr.polls, poll)
}

