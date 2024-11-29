package application

import (
	"errors"

	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
)

type IEnsureColumnOwnership interface {
    Execute(boardId int64, columnId int64) error
}

type ensureColumnOwnership struct {
    repository db.IColumnRepository
}

func NewEnsureColumnOwnership(repository db.IColumnRepository) IEnsureColumnOwnership {
    return &ensureColumnOwnership{
        repository,
    }
}

func (eco *ensureColumnOwnership) Execute(boardId int64, columnId int64) error {
    columns := eco.repository.FindAllByBoardId(boardId)
    found := false
    for _, column := range columns {
        if column.ID == columnId {
            found = true
            break
        }
    }
    if !found {
        return errors.New("COLUMN DOES NOT BELONG TO THE SPECIFIED BOARD")
    }
    return nil
}
