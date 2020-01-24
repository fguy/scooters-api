package scooters

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"github.com/fguy/scooters-api/entities"
	"github.com/fguy/scooters-api/factories"
	mock_repo "github.com/fguy/scooters-api/mocks/repositories/scooters"
	"github.com/golang/mock/gomock"
)

func TestAvailableHandler_ServeHTTP_Success(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock_repo.NewMockInterface(mockCtrl)
	mockRepo.EXPECT().GetAvailable(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(factories.Scooters, nil).Times(1)

	handler := NewAvailableHandler(zap.NewNop(), mockRepo)
	server := httptest.NewServer(handler)
	defer server.Close()

	res, err := server.Client().Get(fmt.Sprint(server.URL, "/api/v1/scooters/available?lat=37.788989&lng=-122.404810&radius=20.0"))
	assert.NoError(t, err)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)

	var scooters []*entities.Scooter
	json.Unmarshal(body, &scooters)
	assert.Equal(t, 2, len(scooters))
}

func TestAvailableHandler_ServeHTTP_Error(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mock_repo.NewMockInterface(mockCtrl)
	mockRepo.EXPECT().GetAvailable(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)

	handler := NewAvailableHandler(zap.NewNop(), mockRepo)
	server := httptest.NewServer(handler)
	defer server.Close()

	res, err := server.Client().Get(fmt.Sprint(server.URL, "/api/v1/scooters/available?lat=37.788989&lng=-122.404810&radius=20.0"))
	assert.NoError(t, err)
	assert.Equal(t, 500, res.StatusCode)
}
