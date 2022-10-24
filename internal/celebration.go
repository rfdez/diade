package diade

import (
	"context"

	"github.com/rfdez/diade/kit/datetime"
	"github.com/rfdez/diade/kit/errors"
	"github.com/rfdez/diade/kit/uuid"
)

// CelebrationID is the type of the ID of a celebration.
type CelebrationID struct {
	value string
}

// NewCelebrationID creates a new CelebrationID.
func NewCelebrationID(value string) (CelebrationID, error) {
	if err := uuid.EnsureIsValidUUID(value); err != nil {
		return CelebrationID{}, errors.NewWrongInput("invalid celebration ID, %s", value)
	}

	return CelebrationID{value: value}, nil
}

// String returns the string representation of the celebration id.
func (id CelebrationID) String() string {
	return id.value
}

// CelebrationDate is the type of the date of a celebration.
type CelebrationDate struct {
	value string
}

// NewCelebrationDate creates a new CelebrationDate.
func NewCelebrationDate(value string) (CelebrationDate, error) {
	if err := datetime.EnsureIsValidDate(value); err != nil {
		return CelebrationDate{}, errors.NewWrongInput("invalid celebration date, %s", value)
	}

	return CelebrationDate{value: value}, nil
}

// String returns the string representation of the celebration date.
func (date CelebrationDate) String() string {
	return date.value
}

// CelebrationName is the type of the name of a celebration.
type CelebrationName struct {
	value string
}

// NewCelebrationName creates a new CelebrationName.
func NewCelebrationName(value string) (CelebrationName, error) {
	if value == "" {
		return CelebrationName{}, errors.NewWrongInput("invalid celebration name, %s", value)
	}

	return CelebrationName{value: value}, nil
}

// String returns the string representation of the celebration name.
func (name CelebrationName) String() string {
	return name.value
}

// CelebrationStatus is the type of the status of a celebration.
type CelebrationStatus struct {
	value string
}

// NewCelebrationStatus creates a new CelebrationStatus.
func NewCelebrationStatus(value string) (CelebrationStatus, error) {
	if value == "" {
		return CelebrationStatus{}, errors.NewWrongInput("invalid celebration status, %s", value)
	}

	return CelebrationStatus{value: value}, nil
}

// String returns the string representation of the celebration status.
func (name CelebrationStatus) String() string {
	return name.value
}

// CelebrationType is the type of the type of a celebration.
type CelebrationType struct {
	value string
}

// NewCelebrationType creates a new CelebrationType.
func NewCelebrationType(value string) (CelebrationType, error) {
	if value == "" {
		return CelebrationType{}, errors.NewWrongInput("invalid celebration type, %s", value)
	}

	return CelebrationType{value: value}, nil
}

// String returns the string representation of the celebration type.
func (name CelebrationType) String() string {
	return name.value
}

// Celebration is the type of a celebration.
type Celebration struct {
	id     CelebrationID
	date   CelebrationDate
	name   CelebrationName
	status CelebrationStatus
	cType  CelebrationType
}

// CelebrationRepository is the interface of the repository of celebrations.
type CelebrationRepository interface {
	// SearchByDate searches for celebrations by date.
	SearchByDate(context.Context, CelebrationDate) ([]Celebration, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CelebrationRepository

// NewCelebration creates a new Celebration.
func NewCelebration(id, date, name, status, t string) (Celebration, error) {
	celebrationID, err := NewCelebrationID(id)
	if err != nil {
		return Celebration{}, err
	}

	celebrationDate, err := NewCelebrationDate(date)
	if err != nil {
		return Celebration{}, err
	}

	celebrationName, err := NewCelebrationName(name)
	if err != nil {
		return Celebration{}, err
	}

	celebrationStatus, err := NewCelebrationStatus(status)
	if err != nil {
		return Celebration{}, err
	}

	celebrationType, err := NewCelebrationType(t)
	if err != nil {
		return Celebration{}, err
	}

	return Celebration{
		id:     celebrationID,
		date:   celebrationDate,
		name:   celebrationName,
		status: celebrationStatus,
		cType:  celebrationType,
	}, nil
}

// ID returns the ID of the celebration.
func (c Celebration) ID() CelebrationID {
	return c.id
}

// Date returns the date of the celebration.
func (c Celebration) Date() CelebrationDate {
	return c.date
}

// Name returns the name of the celebration.
func (c Celebration) Name() CelebrationName {
	return c.name
}

// Status returns the status of the celebration.
func (c Celebration) Status() CelebrationStatus {
	return c.status
}

// Type returns the type of the celebration.
func (c Celebration) Type() CelebrationType {
	return c.cType
}
