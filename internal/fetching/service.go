package fetching

import (
	"context"

	diade "github.com/rfdez/diade/internal"
)

// Service is the interface that provides the fetching methods.
type Service interface {
	// CelebrationsByDate returns the celebrations for the given date.
	CelebrationsByDate(context.Context, string) ([]CelebrationResponse, error)
}

type service struct {
	celebrationRepository diade.CelebrationRepository
}

// NewService creates a new fetching service.
func NewService(celebrationRepository diade.CelebrationRepository) Service {
	return &service{celebrationRepository}
}

// CelebrationsByDate returns the celebrations for the given date.
func (s *service) CelebrationsByDate(ctx context.Context, date string) ([]CelebrationResponse, error) {
	celebrationDate, err := diade.NewCelebrationDate(date)
	if err != nil {
		return nil, err
	}

	celebrations, err := s.celebrationRepository.SearchByDate(ctx, celebrationDate)
	if err != nil {
		return nil, err
	}

	var response []CelebrationResponse
	for _, celebration := range celebrations {
		response = append(response, NewCelebrationResponse(
			celebration.ID().String(),
			celebration.Date().String(),
			celebration.Name().String(),
			celebration.Status().String(),
			celebration.Type().String(),
		))
	}

	return response, nil
}
