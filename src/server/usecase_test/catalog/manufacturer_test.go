package usecase_test

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	mocks "earnforglance/server/domain/mocks"
	test "earnforglance/server/usecase/catalog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestManufacturerUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.ManufacturerRepository)
	timeout := time.Duration(10)
	usecase := test.NewManufacturerUsecase(mockRepo, timeout)

	manufacturerID := bson.NewObjectID().Hex()

	expectedManufacturer := domain.Manufacturer{
		ID:                             bson.NewObjectID(), // Existing ID of the record to update
		Name:                           "Updated TechCorp",
		Description:                    "Updated description for TechCorp",
		ManufacturerID:                 bson.NewObjectID(),
		MetaKeywords:                   "updated, tech, gadgets",
		MetaDescription:                "Updated meta description for TechCorp.",
		MetaTitle:                      "Updated TechCorp - Manufacturer",
		PictureID:                      bson.NewObjectID(),
		PageSize:                       50,
		AllowCustomersToSelectPageSize: false,
		PageSizeOptions:                "50,100",
		SubjectToAcl:                   true,
		LimitedToStores:                true,
		Published:                      false,
		Deleted:                        false,
		DisplayOrder:                   2,
		CreatedOnUtc:                   time.Now().AddDate(0, 0, -30), // Created 30 days ago
		UpdatedOnUtc:                   time.Now(),
		PriceRangeFiltering:            false,
		PriceFrom:                      200.0,
		PriceTo:                        2000.0,
		ManuallyPriceRange:             true,
	}

	mockRepo.On("FetchByID", mock.Anything, manufacturerID).Return(expectedManufacturer, nil)

	result, err := usecase.FetchByID(context.Background(), manufacturerID)

	assert.NoError(t, err)
	assert.Equal(t, expectedManufacturer, result)
	mockRepo.AssertExpectations(t)
}

func TestManufacturerUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.ManufacturerRepository)
	timeout := time.Duration(10)
	usecase := test.NewManufacturerUsecase(mockRepo, timeout)

	newManufacturer := &domain.Manufacturer{
		Name:                           "TechCorp",
		Description:                    "Leading manufacturer of tech products",
		ManufacturerID:                 bson.NewObjectID(),
		MetaKeywords:                   "tech, gadgets, electronics",
		MetaDescription:                "TechCorp is a leading manufacturer of high-quality tech products.",
		MetaTitle:                      "TechCorp - Manufacturer",
		PictureID:                      bson.NewObjectID(),
		PageSize:                       20,
		AllowCustomersToSelectPageSize: true,
		PageSizeOptions:                "10,20,50",
		SubjectToAcl:                   false,
		LimitedToStores:                false,
		Published:                      true,
		Deleted:                        false,
		DisplayOrder:                   1,
		CreatedOnUtc:                   time.Now(),
		UpdatedOnUtc:                   time.Now(),
		PriceRangeFiltering:            true,
		PriceFrom:                      100.0,
		PriceTo:                        1000.0,
		ManuallyPriceRange:             false,
	}

	mockRepo.On("Create", mock.Anything, newManufacturer).Return(nil)

	err := usecase.Create(context.Background(), newManufacturer)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestManufacturerUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.ManufacturerRepository)
	timeout := time.Duration(10)
	usecase := test.NewManufacturerUsecase(mockRepo, timeout)

	updatedManufacturer := &domain.Manufacturer{
		ID:                             bson.NewObjectID(), // Existing ID of the record to update
		Name:                           "Updated TechCorp",
		Description:                    "Updated description for TechCorp",
		ManufacturerID:                 bson.NewObjectID(),
		MetaKeywords:                   "updated, tech, gadgets",
		MetaDescription:                "Updated meta description for TechCorp.",
		MetaTitle:                      "Updated TechCorp - Manufacturer",
		PictureID:                      bson.NewObjectID(),
		PageSize:                       50,
		AllowCustomersToSelectPageSize: false,
		PageSizeOptions:                "50,100",
		SubjectToAcl:                   true,
		LimitedToStores:                true,
		Published:                      false,
		Deleted:                        false,
		DisplayOrder:                   2,
		CreatedOnUtc:                   time.Now().AddDate(0, 0, -30), // Created 30 days ago
		UpdatedOnUtc:                   time.Now(),
		PriceRangeFiltering:            false,
		PriceFrom:                      200.0,
		PriceTo:                        2000.0,
		ManuallyPriceRange:             true,
	}

	mockRepo.On("Update", mock.Anything, updatedManufacturer).Return(nil)

	err := usecase.Update(context.Background(), updatedManufacturer)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestManufacturerUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.ManufacturerRepository)
	timeout := time.Duration(10)
	usecase := test.NewManufacturerUsecase(mockRepo, timeout)

	manufacturerID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, manufacturerID).Return(nil)

	err := usecase.Delete(context.Background(), manufacturerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestManufacturerUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.ManufacturerRepository)
	timeout := time.Duration(10)
	usecase := test.NewManufacturerUsecase(mockRepo, timeout)

	expectedManufacturers := []domain.Manufacturer{
		{
			ID:                             bson.NewObjectID(),
			Name:                           "TechCorp",
			Description:                    "Leading manufacturer of tech products",
			ManufacturerID:                 bson.NewObjectID(),
			MetaKeywords:                   "tech, gadgets, electronics",
			MetaDescription:                "TechCorp is a leading manufacturer of high-quality tech products.",
			MetaTitle:                      "TechCorp - Manufacturer",
			PictureID:                      bson.NewObjectID(),
			PageSize:                       20,
			AllowCustomersToSelectPageSize: true,
			PageSizeOptions:                "10,20,50",
			SubjectToAcl:                   false,
			LimitedToStores:                false,
			Published:                      true,
			Deleted:                        false,
			DisplayOrder:                   1,
			CreatedOnUtc:                   time.Now().AddDate(0, 0, -10), // Created 10 days ago
			UpdatedOnUtc:                   time.Now(),
			PriceRangeFiltering:            true,
			PriceFrom:                      100.0,
			PriceTo:                        1000.0,
			ManuallyPriceRange:             false,
		},
		{
			ID:                             bson.NewObjectID(),
			Name:                           "HomeTech",
			Description:                    "Manufacturer of home appliances",
			ManufacturerID:                 bson.NewObjectID(),
			MetaKeywords:                   "home, appliances, tech",
			MetaDescription:                "HomeTech specializes in high-quality home appliances.",
			MetaTitle:                      "HomeTech - Manufacturer",
			PictureID:                      bson.NewObjectID(),
			PageSize:                       10,
			AllowCustomersToSelectPageSize: false,
			PageSizeOptions:                "10,20",
			SubjectToAcl:                   false,
			LimitedToStores:                false,
			Published:                      true,
			Deleted:                        false,
			DisplayOrder:                   2,
			CreatedOnUtc:                   time.Now().AddDate(0, 0, -20), // Created 20 days ago
			UpdatedOnUtc:                   time.Now(),
			PriceRangeFiltering:            false,
			PriceFrom:                      50.0,
			PriceTo:                        500.0,
			ManuallyPriceRange:             false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedManufacturers, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedManufacturers, result)
	mockRepo.AssertExpectations(t)
}
