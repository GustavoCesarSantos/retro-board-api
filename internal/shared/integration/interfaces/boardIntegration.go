package interfaces

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"

type IBoardApi interface {
	FindAllBoardsByTeamId(teamId int64) ([]domain.Board, error)
	FindAllCardsByColumnId(columnId int64) ([]domain.Card, error)
	FindAllColumnsByBoardId(boardId int64) ([]domain.Column, error)
}