package domain

import (
	"context"
	"time"
)

type Install struct {
	Status       bool
	Details      string
	CreatedOnUtc time.Time
}

type InstallRepository interface {
	PingDatabase(c context.Context) error
}

// GdprLogUsecase interface
type InstallLogUsecase interface {
	PingDatabase(c context.Context) error
}
