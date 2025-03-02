package board

import (
	"go.uber.org/fx"

	boardApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	boardDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/nativeSql"
	boardProvider "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/integration/provider"
	board "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/handlers"
)



var Module = fx.Module(
	"board", 
	fx.Provide(
		//Repositories
		boardDb.NewBoardRepository,
		boardDb.NewCardRepository,
		boardDb.NewColumnRepository,

		//Providers
		boardProvider.NewBoardPublicApiProvider,

		//Applications
		boardApplication.NewFindAllBoards,
		boardApplication.NewFindAllCards,
		boardApplication.NewFindAllColumns,
		boardApplication.NewFindBoard,
		boardApplication.NewFindCard,
		boardApplication.NewFindColumn,
		boardApplication.NewGetNextColumnPosition,
		boardApplication.NewMoveCardBetweenColumns,
		boardApplication.NewMoveColumn,
		boardApplication.NewNotifyMoveCard,
		boardApplication.NewNotifyRemoveCard,
		boardApplication.NewNotifySaveCard,
		boardApplication.NewNotifyUpdateCard,
		boardApplication.NewRemoveBoard,
		boardApplication.NewRemoveCard,
		boardApplication.NewRemoveColumn,
		boardApplication.NewSaveBoard,
		boardApplication.NewSaveCard,
		boardApplication.NewSaveColumn,
		boardApplication.NewUpdateBoard,
		boardApplication.NewUpdateCard,
		boardApplication.NewUpdateColumn,

		//Handlers
		board.NewCreateBoard,
		board.NewCreateCard,
		board.NewCreateColumn,
		board.NewDeleteBoard,
		board.NewDeleteCard,
		board.NewDeleteColumn,
		board.NewEditBoard,
		board.NewEditCard,
		board.NewEditColumn,
		board.NewListAllBoards,
		board.NewListAllCards,
		board.NewListAllColumns,
		board.NewListBoard,
		board.NewListCard,
		board.NewListColumn,
		board.NewMoveCardtoAnotherColumn,
		board.NewMoveColumnToAnotherPosition,

		// Handlers
		board.NewHandlers,
	),
)