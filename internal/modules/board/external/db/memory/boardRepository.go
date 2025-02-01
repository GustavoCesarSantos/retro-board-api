package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type boardRepository struct {
	boards []domain.Board
}

func NewBoardRepository() db.IBoardRepository {
	return &boardRepository{
		boards: []domain.Board{
			*domain.NewBoard(1, 1, "Board 1"),
			*domain.NewBoard(2, 1, "Board 2",),
			*domain.NewBoard(3, 2, "Board 1",),
		},
	}
}

func (br *boardRepository) Delete(boardId int64) error { 
    i := 0
	for _, board := range br.boards {
		if !(board.ID == boardId) {
			br.boards[i] = board
			i++
		}
	}
	br.boards = br.boards[:i]
    return nil
}

func (br *boardRepository) FindAllByTeamId(teamId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Board], error) {
    var boards []domain.Board
    for _, board := range br.boards {
        if board.TeamId == teamId {
            boards = append(boards, board)
        }
    }
	return &utils.ResultPaginated[domain.Board]{
        Items: boards,
        NextCursor: 0,
    }, nil
}

func (br *boardRepository) FindById(boardId int64) (*domain.Board, error) {
    for _, board := range br.boards {
        if board.ID == boardId {
            return &board, nil
        }
    }
    return nil, utils.ErrRecordNotFound
}

func (br *boardRepository) Save(board *domain.Board) error {
	br.boards = append(br.boards, *board)
    return nil
}

func (br *boardRepository) Update(boardId int64, board db.UpdateBoardParams) error {
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
    return nil
}
