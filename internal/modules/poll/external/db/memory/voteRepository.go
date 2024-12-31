package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/external/db/interfaces"
)

type voteRepository struct {
	votes []domain.Vote
}

func NewVoteRepository() db.IVoteRepository {
	return &voteRepository{
		votes: []domain.Vote{
			*domain.NewVote(1, 1, 1),
			*domain.NewVote(1, 2, 1),
			*domain.NewVote(1, 3, 1),
		},
	}
}

func (vr *voteRepository) CountByOptionId(optionId int64) (int, error) {
    var count = 0
    for _, vote := range vr.votes {
        if vote.OptionId == optionId {
            count++
        }
    }
    return count, nil
}

func (vr *voteRepository) Save(vote *domain.Vote) error {
	vr.votes = append(vr.votes, *vote)
	return nil
}

