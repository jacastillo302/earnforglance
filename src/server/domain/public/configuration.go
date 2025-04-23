package domain

import (
	"context"
	domain "earnforglance/server/domain/configuration"
	store "earnforglance/server/domain/stores"
)

type ConfigurationRequest struct {
	ID                        string
	Filters                   []Filter
	Sort                      string
	Limit                     int
	Page                      int
	Lang                      string
	IsRequired                bool
	DisplayDuringRegistration bool
	DisplayOnCustomerInfoPage bool
	Content                   []string
}

type ConfigurationResponse struct {
	Configuration domain.Setting
	Store         *store.Store
}

type ConfigurationsResponse struct {
	Configurations []ConfigurationResponse
}

type ConfigurationRepository interface {
	GetConfigurations(c context.Context, filter ConfigurationRequest) ([]ConfigurationsResponse, error)
}

type ConfigurationUsecase interface {
	GetConfigurations(c context.Context, filter ConfigurationRequest) ([]ConfigurationsResponse, error)
}
