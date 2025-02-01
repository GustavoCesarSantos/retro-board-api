package provider

import (
	"math"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
)

type boardPublicApiProvider struct {
    boardRepository db.IBoardRepository
    cardRepository db.ICardRepository
    columnRepository db.IColumnRepository
}

func NewBoardPublicApiProvider(
    boardRepository db.IBoardRepository,
    cardRepository db.ICardRepository,
    columnRepository db.IColumnRepository,
) interfaces.IBoardApi {
    return &boardPublicApiProvider{
        boardRepository,
        cardRepository,
        columnRepository,
    }
}

func (bpa boardPublicApiProvider) FindAllBoardsByTeamId(teamId int64) ([]domain.Board, error) {
    result, err := bpa.boardRepository.FindAllByTeamId(teamId, math.MaxInt64, math.MaxInt64)
    return result.Items, err
}

func (bpa boardPublicApiProvider) FindAllCardsByColumnId(columnId int64) ([]domain.Card, error) {
    result, err := bpa.cardRepository.FindAllByColumnId(columnId, math.MaxInt64, math.MaxInt64)
    return result.Items, err
}

func (bpa boardPublicApiProvider) FindAllColumnsByBoardId(boardId int64) ([]domain.Column, error) {
    result, err := bpa.columnRepository.FindAllByBoardId(boardId, math.MaxInt64, math.MaxInt64)
    return result.Items, err
}
