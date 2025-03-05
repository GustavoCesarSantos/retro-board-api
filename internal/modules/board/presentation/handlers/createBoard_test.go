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

type MockSaveBoardService struct {
	mock.Mock
}

func (m *MockSaveBoardService) Execute(teamId int64, name string) (*domain.Board, error) {
	args := m.Called(teamId, name)
	return args.Get(0).(*domain.Board), args.Error(1)
}

var createBoardTestCases = []struct {
	name           string
	teamId         string
	input          dtos.CreateBoardRequest
	mockReturn     *domain.Board
	mockError      error
	expectedStatus int
	expectedBody   string
}{
	{
		name:   "Success - Create board with valid input",
		teamId: "1",
		input:  dtos.CreateBoardRequest{Name: "New Board"},
		mockReturn: &domain.Board{
			ID: 1,
			TeamId: 1,
			Name: "New Board",
			CreatedAt: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		},
		mockError:      nil,
		expectedStatus: http.StatusCreated,
		expectedBody:   `{"board":{"id":1,"name":"New Board","created_at":"2023-10-01T00:00:00Z"}}`,
	},
	{
		name:           "Invalid Team ID - Non-numeric team ID",
		teamId:         "invalid",
		input:          dtos.CreateBoardRequest{Name: "New Board"},
		mockReturn:     nil,
		mockError:      nil,
		expectedStatus: http.StatusBadRequest,
		expectedBody:   `{"error":"invalid id parameter"}`,
	},
	{
		name:           "Invalid JSON - Malformed request body",
		teamId:         "1",
		input:          dtos.CreateBoardRequest{},
		mockReturn:     nil,
		mockError:      nil,
		expectedStatus: http.StatusBadRequest,
		expectedBody:   fmt.Sprintf(`{"error": "%s"}`, utils.ErrMissingJSONValue.Error()),
	},
	{
		name:           "Save Board Error - Internal server error",
		teamId:         "1",
		input:          dtos.CreateBoardRequest{Name: "New Board"},
		mockReturn:     nil,
		mockError:      assert.AnError,
		expectedStatus: http.StatusInternalServerError,
		expectedBody:   `{"error":"The server encountered a problem and could not process your request"}`,
	},
}

func TestCreateBoardHandler(t *testing.T) {
	for _, tc := range createBoardTestCases {
		t.Run(tc.name, func(t *testing.T) {
			mockSaveBoard := new(MockSaveBoardService)
			mockSaveBoard.On("Execute", mock.Anything, mock.Anything).Return(tc.mockReturn, tc.mockError)
			handler := &CreateBoard{
				saveBoard: mockSaveBoard,
			}
			inputJSON, _ := json.Marshal(tc.input)
			req := httptest.NewRequest(
                http.MethodPost, 
                fmt.Sprintf("/v1/teams/%s/boards", tc.teamId), 
                bytes.NewBuffer(inputJSON),
            )
			req.Header.Set("Content-Type", "application/json")
            params := httprouter.Params{
                httprouter.Param{Key: "teamId", Value: tc.teamId},
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
