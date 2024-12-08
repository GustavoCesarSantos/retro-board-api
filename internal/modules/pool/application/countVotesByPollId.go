package application

import (
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
)

type Option struct {
	Name  string
	Votes int
}

type Winner struct {
	ID    int64
	Name  string
	Votes int
}

type CountVotesResult struct {
	Options map[int64]Option
	Winner  []Winner
	Total   int
}

type ICountVotesByPollId interface {
	Execute(pollId int64) CountVotesResult
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

func (cv *countVotesByPollId) Execute(pollId int64) CountVotesResult {
	var result CountVotesResult
	result.Options = make(map[int64]Option)
	options := cv.optionRepository.FindAllByPollId(pollId)
	if len(options) == 0 {
		return CountVotesResult{
			Options: make(map[int64]Option),
			Winner:  []Winner{},
			Total:   0,
		}
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
	return result
}
