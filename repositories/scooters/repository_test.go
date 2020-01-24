package scooters

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mock_fx "github.com/fguy/scooters-api/mocks/go.uber.org/fx"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAvailable_Success(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLifecycle := mock_fx.NewMockLifecycle(mockCtrl)
	mockLifecycle.EXPECT().Append(gomock.Any()).Times(1)

	instance, err := New(mockLifecycle, func() (*sql.DB, error) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		rows := sqlmock.NewRows([]string{
			"id", "lat", "lng",
		})
		rows.AddRow(10, 37.788548, -122.411548)
		rows.AddRow(8, 37.783223, -122.398630)
		mock.ExpectQuery(queryAvailable).WillReturnRows(rows).RowsWillBeClosed()

		return db, nil
	})
	assert.NoError(t, err)

	result, err := instance.GetAvailable(
		context.Background(),
		37.788989,
		-122.404810,
		20.0,
	)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.EqualValues(t, 10, result[0].ID)
	assert.EqualValues(t, 8, result[1].ID)
}

func TestGetAvailable_NoRow(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLifecycle := mock_fx.NewMockLifecycle(mockCtrl)
	mockLifecycle.EXPECT().Append(gomock.Any()).Times(1)

	instance, err := New(mockLifecycle, func() (*sql.DB, error) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		rows := sqlmock.NewRows([]string{
			"id", "lat", "lng",
		})
		mock.ExpectQuery(queryAvailable).WillReturnRows(rows).RowsWillBeClosed()

		return db, nil
	})
	assert.NoError(t, err)

	result, err := instance.GetAvailable(
		context.Background(),
		37.788989,
		-122.404810,
		20.0,
	)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetAvailable_QueryError(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLifecycle := mock_fx.NewMockLifecycle(mockCtrl)
	mockLifecycle.EXPECT().Append(gomock.Any()).Times(1)

	instance, err := New(mockLifecycle, func() (*sql.DB, error) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		mock.ExpectQuery(queryAvailable).WillReturnError(errors.New(""))

		return db, nil
	})
	assert.NoError(t, err)

	_, err = instance.GetAvailable(
		context.Background(),
		37.788989,
		-122.404810,
		20.0)
	assert.Error(t, err)
}

func TestGetAvailable_ScanError(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLifecycle := mock_fx.NewMockLifecycle(mockCtrl)
	mockLifecycle.EXPECT().Append(gomock.Any()).Times(1)

	instance, err := New(mockLifecycle, func() (*sql.DB, error) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		rows := sqlmock.NewRows([]string{
			"id", "lat", "lng",
		})
		rows.AddRow(10, 37.788548, -122.411548)
		rows.AddRow("not a number", "37.783223", -122.398630)
		mock.ExpectQuery(queryAvailable).WillReturnRows(rows).RowsWillBeClosed()

		return db, nil
	})
	assert.NoError(t, err)

	_, err = instance.GetAvailable(
		context.Background(),
		37.788989,
		-122.404810,
		20.0,
	)
	assert.Error(t, err)
}

func TestGetAvailable_Error_DB(t *testing.T) {
	t.Parallel()

	_, err := New(nil, func() (*sql.DB, error) {
		return nil, errors.New("")
	})

	assert.Error(t, err)
}

func TestReserve_Success(t *testing.T) {
	t.Parallel()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLifecycle := mock_fx.NewMockLifecycle(mockCtrl)
	mockLifecycle.EXPECT().Append(gomock.Any()).Times(1)

	instance, err := New(mockLifecycle, func() (*sql.DB, error) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectExec(queryReserve).WithArgs(10).WillReturnResult(sqlmock.NewResult(10, 1))

		return db, nil
	})
	assert.NoError(t, err)

	err = instance.Reserve(
		context.Background(),
		10,
	)
	assert.NoError(t, err)
}
