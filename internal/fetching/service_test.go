package fetching_test

import (
	"context"
	"testing"

	"github.com/rfdez/diade/internal/fetching"
	"github.com/rfdez/diade/internal/platform/storage/storagemocks"
	"github.com/rfdez/diade/kit/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_FetchingService_CelebrationsByDate_RepositoryError(t *testing.T) {
	celebrationsRepositoryMock := new(storagemocks.CelebrationRepository)

	celebrationsRepositoryMock.On("SearchByDate", mock.Anything, mock.AnythingOfType("diade.CelebrationDate")).Return(nil, errors.New("error"))

	fetcherService := fetching.NewService(
		celebrationsRepositoryMock,
	)

	_, err := fetcherService.CelebrationsByDate(context.Background(), "2022-01-01")

	celebrationsRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_FetchingService_CelebrationsByDate_InvalidArgumentError(t *testing.T) {
	celebrationsRepositoryMock := new(storagemocks.CelebrationRepository)

	fetcherService := fetching.NewService(
		celebrationsRepositoryMock,
	)

	_, err := fetcherService.CelebrationsByDate(context.Background(), "invalid date")

	assert.Error(t, err)
}

func Test_FetchingService_CelebrationsByDate_Succeed(t *testing.T) {
	celebrationsRepositoryMock := new(storagemocks.CelebrationRepository)

	celebrationsRepositoryMock.On("SearchByDate", mock.Anything, mock.AnythingOfType("diade.CelebrationDate")).Return(nil, nil)

	fetcherService := fetching.NewService(
		celebrationsRepositoryMock,
	)

	_, err := fetcherService.CelebrationsByDate(context.Background(), "2022-01-01")

	celebrationsRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
