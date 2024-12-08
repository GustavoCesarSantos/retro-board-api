package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/domain"

type IVoteRepository interface {
    CountByOptionId(optionId int64) int
	Save(vote domain.Vote)
}

type voteRepository struct {
	votes []domain.Vote
}

func NewVoteRepository() IVoteRepository {
	return &voteRepository{
		votes: []domain.Vote{
			*domain.NewVote(1, 1, 1),
			*domain.NewVote(1, 2, 1),
			*domain.NewVote(1, 3, 1),
		},
	}
}

func (vr *voteRepository) CountByOptionId(optionId int64) int {
    var count = 0
    for _, vote := range vr.votes {
        if vote.OptionId == optionId {
            count++
        }
    }
    return count 
}

func (vr *voteRepository) Save(vote domain.Vote) {
	vr.votes = append(vr.votes, vote)
}

