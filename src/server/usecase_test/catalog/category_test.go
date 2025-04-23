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

func TestCategoryUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewCategoryUsecase(mockRepo, timeout)

	categoryID := bson.NewObjectID().Hex()

	expectedCategory := domain.Category{
		ID:                             bson.NewObjectID(), // Existing ID of the record to update
		Name:                           "Updated Electronics",
		Description:                    "Updated description for electronic products",
		CategoryTemplateID:             bson.NewObjectID(),
		MetaKeywords:                   "updated, electronics",
		MetaDescription:                "Updated meta description for electronics",
		MetaTitle:                      "Updated Electronics",
		ParentCategoryID:               bson.NewObjectID(),
		PictureID:                      bson.NewObjectID(),
		PageSize:                       50,
		AllowCustomersToSelectPageSize: false,
		PageSizeOptions:                "50,100",
		ShowOnHomepage:                 false,
		IncludeInTopMenu:               false,
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
		RestrictFromVendors:            true,
	}

	mockRepo.On("FetchByID", mock.Anything, categoryID).Return(expectedCategory, nil)

	result, err := usecase.FetchByID(context.Background(), categoryID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewCategoryUsecase(mockRepo, timeout)

	newCategory := &domain.Category{
		Name:                           "Electronics",
		Description:                    "Category for electronic products",
		CategoryTemplateID:             bson.NewObjectID(),
		MetaKeywords:                   "electronics, gadgets",
		MetaDescription:                "Find the best electronic products here",
		MetaTitle:                      "Electronics",
		PageSize:                       20,
		AllowCustomersToSelectPageSize: true,
		PageSizeOptions:                "10,20,50",
		ShowOnHomepage:                 true,
		IncludeInTopMenu:               true,
		Published:                      true,
	}

	mockRepo.On("Create", mock.Anything, newCategory).Return(nil)

	err := usecase.Create(context.Background(), newCategory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewCategoryUsecase(mockRepo, timeout)

	updatedCategory := &domain.Category{
		ID:                             bson.NewObjectID(), // Existing ID of the record to update
		Name:                           "Updated Electronics",
		Description:                    "Updated description for electronic products",
		CategoryTemplateID:             bson.NewObjectID(),
		MetaKeywords:                   "updated, electronics",
		MetaDescription:                "Updated meta description for electronics",
		MetaTitle:                      "Updated Electronics",
		ParentCategoryID:               bson.NewObjectID(),
		PictureID:                      bson.NewObjectID(),
		PageSize:                       50,
		AllowCustomersToSelectPageSize: false,
		PageSizeOptions:                "50,100",
		ShowOnHomepage:                 false,
		IncludeInTopMenu:               false,
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
		RestrictFromVendors:            true,
	}

	mockRepo.On("Update", mock.Anything, updatedCategory).Return(nil)

	err := usecase.Update(context.Background(), updatedCategory)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewCategoryUsecase(mockRepo, timeout)

	categoryID := bson.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, categoryID).Return(nil)

	err := usecase.Delete(context.Background(), categoryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.CategoryRepository)
	timeout := time.Duration(10)
	usecase := test.NewCategoryUsecase(mockRepo, timeout)

	expectedCategories := []domain.Category{
		{
			ID:                             bson.NewObjectID(),
			Name:                           "Electronics",
			Description:                    "Category for electronic products",
			CategoryTemplateID:             bson.NewObjectID(),
			MetaKeywords:                   "electronics, gadgets",
			MetaDescription:                "Find the best electronic products here",
			MetaTitle:                      "Electronics",
			ParentCategoryID:               bson.NewObjectID(),
			PictureID:                      bson.NewObjectID(),
			PageSize:                       20,
			AllowCustomersToSelectPageSize: true,
			PageSizeOptions:                "10,20,50",
			ShowOnHomepage:                 true,
			IncludeInTopMenu:               true,
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
			RestrictFromVendors:            false,
		},
		{
			ID:                             bson.NewObjectID(),
			Name:                           "Home Appliances",
			Description:                    "Category for home appliances",
			CategoryTemplateID:             bson.NewObjectID(),
			MetaKeywords:                   "home, appliances",
			MetaDescription:                "Find the best home appliances here",
			MetaTitle:                      "Home Appliances",
			ParentCategoryID:               bson.NewObjectID(),
			PictureID:                      bson.NewObjectID(),
			PageSize:                       10,
			AllowCustomersToSelectPageSize: false,
			PageSizeOptions:                "10,20",
			ShowOnHomepage:                 false,
			IncludeInTopMenu:               true,
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
			RestrictFromVendors:            false,
		},
	}

	mockRepo.On("Fetch", mock.Anything).Return(expectedCategories, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, result)
	mockRepo.AssertExpectations(t)
}
