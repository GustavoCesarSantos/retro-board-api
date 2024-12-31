package provider

import (
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

func (bpa boardPublicApiProvider) FindAllBoardsByTeamId(teamId int64) ([]*domain.Board, error) {
    return bpa.boardRepository.FindAllByTeamId(teamId)
}

func (bpa boardPublicApiProvider) FindAllCardsByColumnId(columnId int64) ([]*domain.Card, error) {
    return bpa.cardRepository.FindAllByColumnId(columnId)
}

func (bpa boardPublicApiProvider) FindAllColumnsByBoardId(boardId int64) ([]*domain.Column, error) {
    return bpa.columnRepository.FindAllByBoardId(boardId)
}
