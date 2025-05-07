package board

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/domain"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSaveCardService struct {
	mock.Mock
}

func (m *MockSaveCardService) Execute(
	columnId int64,
	memberId int64,
	text string,
) (*domain.Card, error) {
	args := m.Called(columnId, memberId, text)
	return args.Get(0).(*domain.Card), args.Error(1)
}

type MockNotifySaveCardService struct {
	mock.Mock
}

func (m *MockNotifySaveCardService) Execute(
	boardId int64,
	columnId int64,
	card *domain.Card,
) error {
	args := m.Called(boardId, columnId, card)
	return args.Error(1)
}

var createCardTestCases = []struct {
	name           string
	teamId         string
	boardId        string
	columnId       string
	input          dtos.CreateCardRequest
	mockReturn     *domain.Card
	mockError      error
	expectedStatus int
	expectedBody   string
}{
	{
		name:     "Success - Create board with valid input",
		teamId:   "1",
		boardId:  "1",
		columnId: "1",
		input:    dtos.CreateCardRequest{Text: "New Card"},
		mockReturn: &domain.Card{
			ID:        1,
			ColumnId:  1,
			MemberId:  1,
			Text:      "New Card",
			CreatedAt: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		},
		mockError:      nil,
		expectedStatus: http.StatusCreated,
		expectedBody:   `{"card":{"id":1,"text":"New Card","created_at":"2023-10-01T00:00:00Z"}}`,
	},
	{
		name:           "Invalid Board ID - Non-numeric team ID",
		teamId:         "1",
		boardId:        "invalid",
		columnId:       "1",
		input:          dtos.CreateCardRequest{Text: "New Card"},
		mockReturn:     nil,
		mockError:      nil,
		expectedStatus: http.StatusBadRequest,
		expectedBody:   `{"error":"invalid id parameter"}`,
	},
	{
		name:           "Invalid Column ID - Non-numeric team ID",
		teamId:         "1",
		boardId:        "1",
		columnId:       "invalid",
		input:          dtos.CreateCardRequest{Text: "New Card"},
		mockReturn:     nil,
		mockError:      nil,
		expectedStatus: http.StatusBadRequest,
		expectedBody:   `{"error":"invalid id parameter"}`,
	},
	{
		name:           "Invalid JSON - Malformed request body",
		teamId:         "1",
		boardId:        "1",
		columnId:       "1",
		input:          dtos.CreateCardRequest{},
		mockReturn:     nil,
		mockError:      nil,
		expectedStatus: http.StatusBadRequest,
		expectedBody:   fmt.Sprintf(`{"error": "%s"}`, utils.ErrMissingJSONValue.Error()),
	},
	{
		name:           "Save Card Error - Internal server error",
		teamId:         "1",
		boardId:        "1",
		columnId:       "1",
		input:          dtos.CreateCardRequest{Text: "New Card"},
		mockReturn:     nil,
		mockError:      assert.AnError,
		expectedStatus: http.StatusInternalServerError,
		expectedBody:   `{"error":"The server encountered a problem and could not process your request"}`,
	},
}

func TestCreateCardHandler(t *testing.T) {
	for _, tc := range createCardTestCases {
		t.Run(tc.name, func(t *testing.T) {
			mockSaveCard := new(MockSaveCardService)
			mockSaveCard.On(
				"Execute",
				mock.Anything,
				mock.Anything,
			).Return(
				tc.mockReturn,
				tc.mockError,
			)
			mockNotifySaveCard := new(MockNotifySaveCardService)
			mockNotifySaveCard.On(
				"Execute",
				mock.Anything,
				mock.Anything,
			).Return(
				tc.mockError,
			)
			handler := &CreateCard{
				notifySaveCard: mockNotifySaveCard,
				saveCard:       mockSaveCard,
			}
			inputJSON, _ := json.Marshal(tc.input)
			req := httptest.NewRequest(
				http.MethodPost,
				fmt.Sprintf("/v1/teams/%s/boards/%s/columns/%s/cards", tc.teamId, tc.boardId, tc.columnId),
				bytes.NewBuffer(inputJSON),
			)
			req.Header.Set("Content-Type", "application/json")
			params := httprouter.Params{
				httprouter.Param{Key: "teamId", Value: tc.teamId},
				httprouter.Param{Key: "boardId", Value: tc.boardId},
				httprouter.Param{Key: "columnId", Value: tc.columnId},
			}
			ctx := context.WithValue(req.Context(), httprouter.ParamsKey, params)
			req = req.WithContext(ctx)
			rec := httptest.NewRecorder()
			handler.Handle(rec, req)
			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedBody, rec.Body.String())
		})
	}
}
