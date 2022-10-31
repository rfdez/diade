package postgresql_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	diade "github.com/rfdez/diade/internal"
	"github.com/rfdez/diade/internal/platform/storage/postgresql"
	"github.com/rfdez/diade/kit/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CelebrationRepository_SearchByDate_RepositoryError(t *testing.T) {
	date, err := diade.NewCelebrationDate("2022-01-01")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		`
		SELECT celebrations.id, celebrations.date, celebrations.name, celebrations.status, celebrations.type FROM celebrations WHERE date = $1`).
		WithArgs(date.String()).
		WillReturnError(errors.New("error"))

	repo := postgresql.NewCelebrationRepository(db, 5*time.Second)

	_, err = repo.SearchByDate(context.Background(), date)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_CelebrationRepository_SearchByDate_Succeed(t *testing.T) {
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	dateVO, err := diade.NewCelebrationDate(date.Format("2006-01-02"))
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "date", "name", "status", "type"}).
		AddRow("2dfb644e-1dbd-4551-ab2c-5d8a2555bcf8", date, "name", "status", "type")

	sqlMock.ExpectQuery(
		`
		SELECT celebrations.id, celebrations.date, celebrations.name, celebrations.status, celebrations.type FROM celebrations WHERE date = $1`).
		WithArgs(dateVO.String()).
		WillReturnRows(rows)

	repo := postgresql.NewCelebrationRepository(db, 5*time.Second)

	_, err = repo.SearchByDate(context.Background(), dateVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
