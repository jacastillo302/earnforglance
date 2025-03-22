package domain

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockAffiliateRepository is a mock implementation of AffiliateRepository
type MockAffiliateRepository struct {
	mock.Mock
}

func (m *MockAffiliateRepository) Create(c context.Context, affiliate *Affiliate) error {
	args := m.Called(c, affiliate)
	return args.Error(0)
}

func (m *MockAffiliateRepository) Update(c context.Context, affiliate *Affiliate) error {
	args := m.Called(c, affiliate)
	return args.Error(0)
}

func (m *MockAffiliateRepository) Delete(c context.Context, affiliate *Affiliate) error {
	args := m.Called(c, affiliate)
	return args.Error(0)
}

func (m *MockAffiliateRepository) Fetch(c context.Context) ([]Affiliate, error) {
	args := m.Called(c)
	return args.Get(0).([]Affiliate), args.Error(1)
}

func (m *MockAffiliateRepository) GetActive(c context.Context, active bool) (Affiliate, error) {
	args := m.Called(c, active)
	return args.Get(0).(Affiliate), args.Error(1)
}

func (m *MockAffiliateRepository) FetchByID(c context.Context, affiliateID string) (Affiliate, error) {
	args := m.Called(c, affiliateID)
	return args.Get(0).(Affiliate), args.Error(1)
}
