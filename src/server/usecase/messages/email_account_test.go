package usecase

import (
	"context"
	domain "earnforglance/server/domain/messages"
	mocks "earnforglance/server/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestEmailAccountUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.EmailAccountRepository)
	timeout := time.Duration(10)
	usecase := NewEmailAccountUsecase(mockRepo, timeout)

	emailAccountID := primitive.NewObjectID().Hex()

	updatedEmailAccount := domain.EmailAccount{
		ID:                          primitive.NewObjectID(), // Existing ID of the record to update
		Email:                       "updated@example.com",
		DisplayName:                 "Updated Account",
		Host:                        "smtp.updated.com",
		Port:                        465,
		Username:                    "updated_user",
		Password:                    "updatedpassword",
		EnableSsl:                   false,
		MaxNumberOfEmails:           200,
		EmailAuthenticationMethodID: 2,
		ClientID:                    "updated-client-id",
		ClientSecret:                "updated-client-secret",
		TenantID:                    "updated-tenant-id",
		EmailAuthenticationMethod:   1,
	}

	mockRepo.On("FetchByID", mock.Anything, emailAccountID).Return(updatedEmailAccount, nil)

	result, err := usecase.FetchByID(context.Background(), emailAccountID)

	assert.NoError(t, err)
	assert.Equal(t, updatedEmailAccount, result)
	mockRepo.AssertExpectations(t)
}

func TestEmailAccountUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.EmailAccountRepository)
	timeout := time.Duration(10)
	usecase := NewEmailAccountUsecase(mockRepo, timeout)

	newEmailAccount := &domain.EmailAccount{
		Email:                       "example@example.com",
		DisplayName:                 "Example Account",
		Host:                        "smtp.example.com",
		Port:                        587,
		Username:                    "example_user",
		Password:                    "securepassword",
		EnableSsl:                   true,
		MaxNumberOfEmails:           100,
		EmailAuthenticationMethodID: 1,
		ClientID:                    "client-id-example",
		ClientSecret:                "client-secret-example",
		TenantID:                    "tenant-id-example",
		EmailAuthenticationMethod:   2,
	}

	mockRepo.On("Create", mock.Anything, newEmailAccount).Return(nil)

	err := usecase.Create(context.Background(), newEmailAccount)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEmailAccountUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.EmailAccountRepository)
	timeout := time.Duration(10)
	usecase := NewEmailAccountUsecase(mockRepo, timeout)

	updatedEmailAccount := &domain.EmailAccount{
		ID:                          primitive.NewObjectID(), // Existing ID of the record to update
		Email:                       "updated@example.com",
		DisplayName:                 "Updated Account",
		Host:                        "smtp.updated.com",
		Port:                        465,
		Username:                    "updated_user",
		Password:                    "updatedpassword",
		EnableSsl:                   false,
		MaxNumberOfEmails:           200,
		EmailAuthenticationMethodID: 2,
		ClientID:                    "updated-client-id",
		ClientSecret:                "updated-client-secret",
		TenantID:                    "updated-tenant-id",
		EmailAuthenticationMethod:   1,
	}

	mockRepo.On("Update", mock.Anything, updatedEmailAccount).Return(nil)

	err := usecase.Update(context.Background(), updatedEmailAccount)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEmailAccountUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.EmailAccountRepository)
	timeout := time.Duration(10)
	usecase := NewEmailAccountUsecase(mockRepo, timeout)

	emailAccountID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, emailAccountID).Return(nil)

	err := usecase.Delete(context.Background(), emailAccountID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEmailAccountUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.EmailAccountRepository)
	timeout := time.Duration(10)
	usecase := NewEmailAccountUsecase(mockRepo, timeout)
	fetchedEmailAccounts := []domain.EmailAccount{
		{
			ID:                          primitive.NewObjectID(),
			Email:                       "example@example.com",
			DisplayName:                 "Example Account",
			Host:                        "smtp.example.com",
			Port:                        587,
			Username:                    "example_user",
			Password:                    "securepassword",
			EnableSsl:                   true,
			MaxNumberOfEmails:           100,
			EmailAuthenticationMethodID: 1,
			ClientID:                    "client-id-example",
			ClientSecret:                "client-secret-example",
			TenantID:                    "tenant-id-example",
			EmailAuthenticationMethod:   2,
		},
		{
			ID:                          primitive.NewObjectID(),
			Email:                       "updated@example.com",
			DisplayName:                 "Updated Account",
			Host:                        "smtp.updated.com",
			Port:                        465,
			Username:                    "updated_user",
			Password:                    "updatedpassword",
			EnableSsl:                   false,
			MaxNumberOfEmails:           200,
			EmailAuthenticationMethodID: 2,
			ClientID:                    "updated-client-id",
			ClientSecret:                "updated-client-secret",
			TenantID:                    "updated-tenant-id",
			EmailAuthenticationMethod:   0,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(fetchedEmailAccounts, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedEmailAccounts, result)
	mockRepo.AssertExpectations(t)
}
