package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"

type UpdateBoardParams struct {
	Name *string
	Active *bool
}

type IBoardRepository interface {
	Delete(boardId int64)
    FindAllByTeamId(teamId int64) []*domain.Board
    FindById(boardId int64) *domain.Board
	Save(board domain.Board)
	Update(boardId int64, board UpdateBoardParams)
}

type boardRepository struct {
	boards []domain.Board
}

func NewBoardRepository() IBoardRepository {
	return &boardRepository{
		boards: []domain.Board{
			*domain.NewBoard(1, 1, "Board 1"),
			*domain.NewBoard(2, 1, "Board 2",),
			*domain.NewBoard(3, 2, "Board 1",),
		},
	}
}

func (br *boardRepository) Delete(boardId int64) {
    i := 0
	for _, board := range br.boards {
		if !(board.ID == boardId) {
			br.boards[i] = board
			i++
		}
	}
	br.boards = br.boards[:i]
}

func (br *boardRepository) FindAllByTeamId(teamId int64) []*domain.Board {
    var boards []*domain.Board
    for _, board := range br.boards {
        if board.TeamId == teamId {
            boards = append(boards, &board)
        }
    }
    return boards
}

func (br *boardRepository) FindById(boardId int64) *domain.Board {
    for _, board := range br.boards {
        if board.ID == boardId {
            return &board
        }
    }
    return nil
}

func (br *boardRepository) Save(board domain.Board) {
	br.boards = append(br.boards, board)
}

func (br *boardRepository) Update(boardId int64, board UpdateBoardParams) {
    for i := range br.boards {
		if br.boards[i].ID == boardId {
			if board.Name != nil {
				br.boards[i].Name = *board.Name
			}
			if board.Active != nil {
				br.boards[i].Active = *board.Active
			}
			break
		}
	}
}
