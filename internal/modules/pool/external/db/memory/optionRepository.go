package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/domain"

type IOptionRepository interface {
	Delete(optionId int64)
    FindAllByPollId(pollId int64) []*domain.Option
	Save(option domain.Option)
}

type optionRepository struct {
	options []domain.Option
}

func NewOptionRepository() IOptionRepository {
	return &optionRepository{
		options: []domain.Option{
			*domain.NewOption(1, 1, "Option 1"),
			*domain.NewOption(2, 1, "Option 2",),
			*domain.NewOption(3, 2, "Option 1",),
		},
	}
}

func (or *optionRepository) Delete(optionId int64) {
    i := 0
	for _, option := range or.options {
		if !(option.ID == optionId) {
			or.options[i] = option
			i++
		}
	}
	or.options = or.options[:i]
}

func (or *optionRepository) FindAllByPollId(pollId int64) []*domain.Option {
    var options []*domain.Option
    for _, option := range or.options {
        if option.PollId == pollId {
            options = append(options, &option)
        }
    }
    return options
}

func (or *optionRepository) Save(option domain.Option) {
	or.options = append(or.options, option)
}

