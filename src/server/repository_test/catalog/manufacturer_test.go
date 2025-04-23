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

type MockSingleResultManufacturer struct {
	mock.Mock
}

func (m *MockSingleResultManufacturer) Decode(v interface{}) error {
	args := m.Called(v)
	if result, ok := args.Get(0).(*domain.Manufacturer); ok {
		*v.(*domain.Manufacturer) = *result
	}
	return args.Error(1)
}

func TestManufacturerRepository_FetchByID(t *testing.T) {
	var databaseHelper *mocks.Database
	var collectionHelper *mocks.Collection

	databaseHelper = &mocks.Database{}
	collectionHelper = &mocks.Collection{}

	collectionName := domain.CollectionManufacturer

	mockItem := domain.Manufacturer{ID: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, Name: "", Description: "", ManufacturerID: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, MetaKeywords: "", MetaDescription: "", MetaTitle: "", PictureID: bson.ObjectID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, PageSize: 0, AllowCustomersToSelectPageSize: false, PageSizeOptions: "", SubjectToAcl: false, LimitedToStores: false, Published: false, Deleted: false, DisplayOrder: 0, CreatedOnUtc: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedOnUtc: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), PriceRangeFiltering: false, PriceFrom: 0, PriceTo: 0, ManuallyPriceRange: false}

	t.Run("success", func(t *testing.T) {
		mockSingleResult := &MockSingleResultManufacturer{}
		mockSingleResult.On("Decode", mock.Anything).Return(mockItem, nil)

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewManufacturerRepository(databaseHelper, collectionName)

		result, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, mockItem, result)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockSingleResult := &MockSingleResultManufacturer{}
		mockSingleResult.On("Decode", mock.Anything).Return(nil, errors.New("Unexpected"))

		collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult).Once()
		databaseHelper.On("Collection", collectionName).Return(collectionHelper)

		ur := repository.NewManufacturerRepository(databaseHelper, collectionName)

		_, err := ur.FetchByID(context.Background(), mockItem.ID.Hex())

		assert.Error(t, err)

		collectionHelper.AssertExpectations(t)
		mockSingleResult.AssertExpectations(t)
	})
}

func TestManufacturerRepository_Create(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionManufacturer

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockManufacturer := &domain.Manufacturer{
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

	collectionHelper.On("InsertOne", mock.Anything, mockManufacturer).Return(nil, nil).Once()

	repo := repository.NewManufacturerRepository(databaseHelper, collectionName)

	err := repo.Create(context.Background(), mockManufacturer)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}

func TestManufacturerRepository_Update(t *testing.T) {
	databaseHelper := &mocks.Database{}
	collectionHelper := &mocks.Collection{}
	collectionName := domain.CollectionManufacturer

	databaseHelper.On("Collection", collectionName).Return(collectionHelper)

	mockManufacturer := &domain.Manufacturer{
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

	filter := bson.M{"_id": mockManufacturer.ID}
	update := bson.M{"$set": mockManufacturer}

	collectionHelper.On("UpdateOne", mock.Anything, filter, update).Return(nil, nil).Once()

	repo := repository.NewManufacturerRepository(databaseHelper, collectionName)

	err := repo.Update(context.Background(), mockManufacturer)

	assert.NoError(t, err)
	collectionHelper.AssertExpectations(t)
}
