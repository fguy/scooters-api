package scooters

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	mock_repo "github.com/fguy/scooters-api/mocks/repositories/scooters"
	"github.com/golang/mock/gomock"
)

func TestReserveHandler_ServeHTTP_Success(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock_repo.NewMockInterface(mockCtrl)
	mockRepo.EXPECT().Reserve(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	handler := NewReserveHandler(zap.NewNop(), mockRepo)
	server := httptest.NewServer(handler)
	defer server.Close()

	_, err := server.Client().Get(fmt.Sprint(server.URL, "/api/v1/scooters/reserve?id=10"))
	assert.NoError(t, err)
}

func TestReserveHandler_ServeHTTP_Error(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock_repo.NewMockInterface(mockCtrl)
	mockRepo.EXPECT().Reserve(gomock.Any(), gomock.Any()).Return(errors.New("")).Times(1)

	handler := NewReserveHandler(zap.NewNop(), mockRepo)
	server := httptest.NewServer(handler)
	defer server.Close()

	res, err := server.Client().Get(fmt.Sprint(server.URL, "/api/v1/scooters/reserve?id=10"))
	assert.NoError(t, err)
	assert.Equal(t, 500, res.StatusCode)
}
