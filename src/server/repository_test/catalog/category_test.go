package repository_test

import (
	"context"
	domain "earnforglance/server/domain/catalog"
	repository "earnforglance/server/repository/catalog"
	"earnforglance/server/service/data/mongo/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockSingleResultCategory struct {
	mock.Mock
}

func (m *MockSingleResultCategory) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Category); ok {
		*v.(*domain.Category) = *result
	}
	return args.Error(1)
}

func TestCategoryRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionCategory

	mockItem := domain.Category{ID: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, Name: "", Description: "", CategoryTemplateID: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, MetaKeywords: "", MetaDescription: "", MetaTitle: "", ParentCategoryID: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, PictureID: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, PageSize: 0, AllowCustomersToSelectPageSize: false, PageSizeOptions: "", ShowOnHomepage: false, IncludeInTopMenu: false, SubjectToAcl: false, LimitedToStores: false, Published: false, Deleted: false, DisplayOrder: 0, CreatedOnUtc: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedOnUtc: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), PriceRangeFiltering: false, PriceFrom: 0, PriceTo: 0, ManuallyPriceRange: false, RestrictFromVendors: false}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCategory{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCategoryRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultCategory{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewCategoryRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestCategoryRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCategory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockCategory := &domain.Category{
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

	collectionHelper.On("InsertOne", mock.Anything, mockCategory).Return(nil, nil).Once()

	repo := repository.NewCategoryRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockCategory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestCategoryRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionCategory

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockCategory := &domain.Category{
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

	filter := bson.M{"_id": mockCategory.ID}
	update := bson.M{"$set": mockCategory}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewCategoryRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockCategory)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
