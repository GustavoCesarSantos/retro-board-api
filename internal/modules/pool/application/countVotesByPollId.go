package application

import (
	"errors"

	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
)

type Option struct {
	Name  string `json:"name" example:"1"`
	Votes int `json:"votes" example:"1"`
}

type Winner struct {
	ID    int64 `json:"id" example:"1"`
	Name  string `json:"name" example:"1"`
	Votes int `json:"votes" example:"1"`
}

type CountVotesResult struct {
	Options map[int64]Option `json:"options"`
	Winner  []Winner `json:"winner"`
	Total   int `json:"total"`
}

type ICountVotesByPollId interface {
	Execute(pollId int64) (*CountVotesResult, error)
}

type countVotesByPollId struct {
	optionRepository db.IOptionRepository
	voteRepository   db.IVoteRepository
}

func NewCountVotesByPollId(
	optionRepository db.IOptionRepository,
	voteRepository db.IVoteRepository,
) ICountVotesByPollId {
	return &countVotesByPollId{
		optionRepository,
		voteRepository,
	}
}

func (cv *countVotesByPollId) Execute(pollId int64) (*CountVotesResult, error) {
	var result CountVotesResult
	result.Options = make(map[int64]Option)
	options := cv.optionRepository.FindAllByPollId(pollId)
	if len(options) == 0 {
		return nil, errors.New("FAILURE TO COUNT VOTES")	
    }
	for _, option := range options {
		count := cv.voteRepository.CountByOptionId(option.ID)
		result.Total += count
        result.Options[option.ID] = Option{
			Name: option.Text,
			Votes: count,
		}
		if len(result.Winner) == 0 || count > result.Winner[0].Votes {
			result.Winner = []Winner{{
				ID:    option.ID,
				Name:  option.Text,
				Votes: count,
			}}
		} else if count == result.Winner[0].Votes {
			result.Winner = append(result.Winner, Winner{
				ID:    option.ID,
				Name:  option.Text,
				Votes: count,
			})
		}
	}
	return &result, nil
}
