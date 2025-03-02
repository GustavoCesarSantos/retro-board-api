package board

type Handlers struct {
	CreateBoard *CreateBoard
	CreateCard *CreateCard
	CreateColumn *CreateColumn
	DeleteBoard *DeleteBoard
	DeleteCard *DeleteCard
	DeleteColumn *DeleteColumn
	EditBoard *EditBoard
	EditCard *EditCard
	EditColumn *EditColumn
	ListAllBoards *ListAllBoards
	ListAllCards *ListAllCards
	ListAllColumns *ListAllColumns
	ListBoard *ListBoard
	ListCard *ListCard
	ListColumn *ListColumn
	MoveCardToAnotherColumn *MoveCardToAnotherColumn
	MoveColumnToAnotherPosition *MoveColumnToAnotherPosition
}

func NewHandlers(
	createBoard *CreateBoard,
	createCard *CreateCard,
	createColumn *CreateColumn,
	deleteBoard *DeleteBoard,
	deleteCard *DeleteCard,
	deleteColumn *DeleteColumn,
	editBoard *EditBoard,
	editCard *EditCard,
	editColumn *EditColumn,
	listAllBoards *ListAllBoards,
	listAllCards *ListAllCards,
	listAllColumns *ListAllColumns,
	listBoard *ListBoard,
	listCard *ListCard,
	listColumn *ListColumn,
	moveCardToAnotherColumn *MoveCardToAnotherColumn,
	moveColumnToAnotherPosition *MoveColumnToAnotherPosition,
) *Handlers {
	return &Handlers{
		CreateBoard: createBoard,
		CreateCard: createCard,
		CreateColumn: createColumn,
		DeleteBoard: deleteBoard,
		DeleteCard: deleteCard,
		DeleteColumn: deleteColumn,
		EditBoard: editBoard,
		EditCard: editCard,
		EditColumn: editColumn,
		ListAllBoards: listAllBoards,
		ListAllCards: listAllCards,
		ListAllColumns: listAllColumns,
		ListBoard: listBoard,
		ListCard: listCard,
		ListColumn: listColumn,
		MoveCardToAnotherColumn: moveCardToAnotherColumn,
		MoveColumnToAnotherPosition: moveColumnToAnotherPosition,
	}
}