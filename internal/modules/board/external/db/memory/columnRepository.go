package db

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)



type columnRepository struct {
	columns []domain.Column
}

func NewColumnRepository() db.IColumnRepository {
	return &columnRepository{
		columns: []domain.Column{
			*domain.NewColumn(1, 1, "Column 1", "#FFFFFF", 1),
			*domain.NewColumn(2, 1, "Column 2", "#FFFFFG", 2),
			*domain.NewColumn(3, 1, "Column 3", "#FFFFFH", 3),
		},
	}
}

func (cr *columnRepository) CountColumnsByBoardId(boardId int64) (int, error) {
    return len(cr.columns), nil
}

func (cr *columnRepository) Delete(columnId int64) error {
    i := 0
	for _, column := range cr.columns {
		if !(column.ID == columnId) {
			cr.columns[i] = column
			i++
		}
	}
	cr.columns = cr.columns[:i]
    return nil
}

func (cr *columnRepository) FindAllByBoardId(boardId int64, limit int, lastId int) (*utils.ResultPaginated[domain.Column], error) {
    var columns []domain.Column
    for _, column := range cr.columns {
        if column.BoardId == boardId {
            columns = append(columns, column)
        }
    }
	return &utils.ResultPaginated[domain.Column]{
        Items: columns,
        NextCursor: 0,
    }, nil
}

func (cr *columnRepository) Save(column *domain.Column) error {
	cr.columns = append(cr.columns, *column)
    return nil
}

func (cr *columnRepository) Update(columnId int64, column db.UpdateColumnParams) error {
    for i := range cr.columns {
		if cr.columns[i].ID == columnId {
			if column.Name != nil {
				cr.columns[i].Name = *column.Name
			}
			if column.Color != nil {
				cr.columns[i].Color = *column.Color
			}
			break
		}
	}
    return nil
}
