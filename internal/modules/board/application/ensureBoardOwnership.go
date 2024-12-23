package application

import (
	"errors"

	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
)

type IEnsureBoardOwnership interface {
    Execute(teamId int64, boardId int64) error
}

type ensureBoardOwnership struct {
    repository db.IBoardRepository
}

func NewEnsureBoardOwnership(repository db.IBoardRepository) IEnsureBoardOwnership {
    return &ensureBoardOwnership{
        repository,
    }
}

func (ebo *ensureBoardOwnership) Execute(teamId int64, boardId int64) error {
    boards, findErr := ebo.repository.FindAllByTeamId(teamId)
    if findErr != nil {
        return findErr
    }
    found := false
    for _, board := range boards {
        if board.ID == boardId {
            found = true
            break
        }
    }
    if !found {
        return errors.New("BOARD DOES NOT BELONG TO THE SPECIFIED TEAM")
    }
    return nil
}
