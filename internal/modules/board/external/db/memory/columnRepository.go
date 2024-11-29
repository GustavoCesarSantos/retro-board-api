package db

import "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"

type UpdateColumnParams struct {
	Name *string
    Color *string
}

type IColumnRepository interface {
	CountColumnsByBoardId(boardId int64) int
	Delete(columnId int64)
    FindAllByBoardId(boardId int64) []*domain.Column
	Save(column domain.Column)
	Update(columnId int64, column UpdateColumnParams)
}

type columnRepository struct {
	columns []domain.Column
}

func NewColumnRepository() IColumnRepository {
	return &columnRepository{
		columns: []domain.Column{
			*domain.NewColumn(1, 1, "Column 1", "#FFFFFF", 1),
			*domain.NewColumn(2, 1, "Column 2", "#FFFFFG", 2),
			*domain.NewColumn(3, 1, "Column 3", "#FFFFFH", 3),
		},
	}
}

func (cr *columnRepository) CountColumnsByBoardId(boardId int64) int {
    return len(cr.columns)
}

func (cr *columnRepository) Delete(columnId int64) {
    i := 0
	for _, column := range cr.columns {
		if !(column.ID == columnId) {
			cr.columns[i] = column
			i++
		}
	}
	cr.columns = cr.columns[:i]
}

func (cr *columnRepository) FindAllByBoardId(boardId int64) []*domain.Column {
    var columns []*domain.Column
    for _, column := range cr.columns {
        if column.BoardId == boardId {
            columns = append(columns, &column)
        }
    }
    return columns
}

func (cr *columnRepository) Save(column domain.Column) {
	cr.columns = append(cr.columns, column)
}

func (cr *columnRepository) Update(columnId int64, column UpdateColumnParams) {
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
}
