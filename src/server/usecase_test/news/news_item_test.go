package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/news"
	test "earnforglance/server/usecase/news"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewsItemUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.NewsItemRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewNewsItemUsecase(mockRepo, timeout)

	newsItemID := primitive.NewObjectID().Hex()

	updatedNewsItem := domain.NewsItem{
		ID:              primitive.NewObjectID(), // Existing ID of the record to update
		LanguageID:      primitive.NewObjectID(),
		Title:           "Updated Feature Announcement",
		Short:           "We have updated the feature announcement.",
		Full:            "The new feature has been updated to include additional functionality.",
		Published:       false,
		StartDateUtc:    new(time.Time),
		EndDateUtc:      new(time.Time),
		AllowComments:   false,
		LimitedToStores: true,
		MetaKeywords:    "feature, update",
		MetaDescription: "Updated announcement of a feature.",
		MetaTitle:       "Updated Feature Announcement",
		CreatedOnUtc:    time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}

	mockRepo.On("FetchByID", mock.Anything, newsItemID).Return(updatedNewsItem, nil)

	result, err := usecase.FetchByID(context.Background(), newsItemID)

	assert.NoError(t, err)
	assert.Equal(t, updatedNewsItem, result)
	mockRepo.AssertExpectations(t)
}

func TestNewsItemUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.NewsItemRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewNewsItemUsecase(mockRepo, timeout)

	newNewsItem := &domain.NewsItem{
		LanguageID:      primitive.NewObjectID(),
		Title:           "New Feature Announcement",
		Short:           "We are excited to announce a new feature.",
		Full:            "Our platform now supports a new feature that will enhance your experience.",
		Published:       true,
		StartDateUtc:    nil,
		EndDateUtc:      nil,
		AllowComments:   true,
		LimitedToStores: false,
		MetaKeywords:    "feature, announcement",
		MetaDescription: "Announcement of a new feature.",
		MetaTitle:       "New Feature Announcement",
		CreatedOnUtc:    time.Now(),
	}
	mockRepo.On("Create", mock.Anything, newNewsItem).Return(nil)

	err := usecase.Create(context.Background(), newNewsItem)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsItemUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.NewsItemRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewNewsItemUsecase(mockRepo, timeout)

	updatedNewsItem := &domain.NewsItem{
		ID:              primitive.NewObjectID(), // Existing ID of the record to update
		LanguageID:      primitive.NewObjectID(),
		Title:           "Updated Feature Announcement",
		Short:           "We have updated the feature announcement.",
		Full:            "The new feature has been updated to include additional functionality.",
		Published:       false,
		StartDateUtc:    new(time.Time),
		EndDateUtc:      new(time.Time),
		AllowComments:   false,
		LimitedToStores: true,
		MetaKeywords:    "feature, update",
		MetaDescription: "Updated announcement of a feature.",
		MetaTitle:       "Updated Feature Announcement",
		CreatedOnUtc:    time.Now().AddDate(0, 0, -7), // Created 7 days ago
	}
	*updatedNewsItem.StartDateUtc = time.Now().AddDate(0, 0, 1) // Start date is tomorrow
	*updatedNewsItem.EndDateUtc = time.Now().AddDate(0, 0, 10)  // End date is in 10 days

	mockRepo.On("Update", mock.Anything, updatedNewsItem).Return(nil)

	err := usecase.Update(context.Background(), updatedNewsItem)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsItemUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.NewsItemRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewNewsItemUsecase(mockRepo, timeout)

	newsItemID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, newsItemID).Return(nil)

	err := usecase.Delete(context.Background(), newsItemID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestNewsItemUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.NewsItemRepository)
	timeout := time.Duration(10) * time.Second
	usecase := test.NewNewsItemUsecase(mockRepo, timeout)
	fetchedNewsItems := []domain.NewsItem{
		{
			ID:              primitive.NewObjectID(),
			LanguageID:      primitive.NewObjectID(),
			Title:           "New Feature Announcement",
			Short:           "We are excited to announce a new feature.",
			Full:            "Our platform now supports a new feature that will enhance your experience.",
			Published:       true,
			StartDateUtc:    nil,
			EndDateUtc:      nil,
			AllowComments:   true,
			LimitedToStores: false,
			MetaKeywords:    "feature, announcement",
			MetaDescription: "Announcement of a new feature.",
			MetaTitle:       "New Feature Announcement",
			CreatedOnUtc:    time.Now().AddDate(0, 0, -10), // Created 10 days ago
		},
		{
			ID:              primitive.NewObjectID(),
			LanguageID:      primitive.NewObjectID(),
			Title:           "Maintenance Update",
			Short:           "Scheduled maintenance update.",
			Full:            "Our platform will undergo scheduled maintenance to improve performance.",
			Published:       false,
			StartDateUtc:    new(time.Time),
			EndDateUtc:      new(time.Time),
			AllowComments:   false,
			LimitedToStores: true,
			MetaKeywords:    "maintenance, update",
			MetaDescription: "Scheduled maintenance update.",
			MetaTitle:       "Maintenance Update",
			CreatedOnUtc:    time.Now().AddDate(0, 0, -5), // Created 5 days ago
		},
	}
	*fetchedNewsItems[1].StartDateUtc = time.Now().AddDate(0, 0, 2) // Start date is in 2 days
	*fetchedNewsItems[1].EndDateUtc = time.Now().AddDate(0, 0, 7)   // End date is in 7 days

	mockRepo.On("Fetch", mock.Anything).Return(fetchedNewsItems, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedNewsItems, result)
	mockRepo.AssertExpectations(t)
}
