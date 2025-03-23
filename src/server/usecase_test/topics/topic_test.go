package usecase_test

import (
	"context"
	mocks "earnforglance/server/domain/mocks"
	domain "earnforglance/server/domain/topics"
	test "earnforglance/server/usecase/topics"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTopicUsecase_FetchByID(t *testing.T) {
	mockRepo := new(mocks.TopicRepository)
	timeout := time.Duration(10)
	usecase := test.NewTopicUsecase(mockRepo, timeout)

	topicID := primitive.NewObjectID().Hex()

	updatedTopic := domain.Topic{
		ID:                        primitive.NewObjectID(), // Existing ID of the record to update
		SystemName:                "contact-us",
		IncludeInSitemap:          true,
		IncludeInTopMenu:          false,
		IncludeInFooterColumn1:    true,
		IncludeInFooterColumn2:    false,
		IncludeInFooterColumn3:    true,
		DisplayOrder:              2,
		AccessibleWhenStoreClosed: false,
		IsPasswordProtected:       true,
		Password:                  "securepassword",
		Title:                     "Contact Us",
		Body:                      "This is the Contact Us page content.",
		Published:                 true,
		TopicTemplateID:           primitive.NewObjectID(),
		MetaKeywords:              "contact, support, help",
		MetaDescription:           "Get in touch with us for support.",
		MetaTitle:                 "Contact Us - Support",
		SubjectToAcl:              true,
		LimitedToStores:           true,
		AvailableStartDateTimeUtc: new(time.Time),
		AvailableEndDateTimeUtc:   new(time.Time),
	}

	mockRepo.On("FetchByID", mock.Anything, topicID).Return(updatedTopic, nil)

	result, err := usecase.FetchByID(context.Background(), topicID)

	assert.NoError(t, err)
	assert.Equal(t, updatedTopic, result)
	mockRepo.AssertExpectations(t)
}

func TestTopicUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.TopicRepository)
	timeout := time.Duration(10)
	usecase := test.NewTopicUsecase(mockRepo, timeout)

	newTopic := &domain.Topic{
		SystemName:                "about-us",
		IncludeInSitemap:          true,
		IncludeInTopMenu:          true,
		IncludeInFooterColumn1:    false,
		IncludeInFooterColumn2:    true,
		IncludeInFooterColumn3:    false,
		DisplayOrder:              1,
		AccessibleWhenStoreClosed: true,
		IsPasswordProtected:       false,
		Password:                  "",
		Title:                     "About Us",
		Body:                      "This is the About Us page content.",
		Published:                 true,
		TopicTemplateID:           primitive.NewObjectID(),
		MetaKeywords:              "about, company, info",
		MetaDescription:           "Learn more about our company.",
		MetaTitle:                 "About Us - Company Info",
		SubjectToAcl:              false,
		LimitedToStores:           false,
		AvailableStartDateTimeUtc: nil,
		AvailableEndDateTimeUtc:   nil,
	}

	mockRepo.On("Create", mock.Anything, newTopic).Return(nil)

	err := usecase.Create(context.Background(), newTopic)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTopicUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.TopicRepository)
	timeout := time.Duration(10)
	usecase := test.NewTopicUsecase(mockRepo, timeout)

	updatedTopic := &domain.Topic{
		ID:                        primitive.NewObjectID(), // Existing ID of the record to update
		SystemName:                "contact-us",
		IncludeInSitemap:          true,
		IncludeInTopMenu:          false,
		IncludeInFooterColumn1:    true,
		IncludeInFooterColumn2:    false,
		IncludeInFooterColumn3:    true,
		DisplayOrder:              2,
		AccessibleWhenStoreClosed: false,
		IsPasswordProtected:       true,
		Password:                  "securepassword",
		Title:                     "Contact Us",
		Body:                      "This is the Contact Us page content.",
		Published:                 true,
		TopicTemplateID:           primitive.NewObjectID(),
		MetaKeywords:              "contact, support, help",
		MetaDescription:           "Get in touch with us for support.",
		MetaTitle:                 "Contact Us - Support",
		SubjectToAcl:              true,
		LimitedToStores:           true,
		AvailableStartDateTimeUtc: new(time.Time),
		AvailableEndDateTimeUtc:   new(time.Time),
	}
	*updatedTopic.AvailableStartDateTimeUtc = time.Now().AddDate(0, 0, -7) // Available from 7 days ago
	*updatedTopic.AvailableEndDateTimeUtc = time.Now().AddDate(0, 0, 7)    // Available for the next 7 days

	mockRepo.On("Update", mock.Anything, updatedTopic).Return(nil)

	err := usecase.Update(context.Background(), updatedTopic)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTopicUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.TopicRepository)
	timeout := time.Duration(10)
	usecase := test.NewTopicUsecase(mockRepo, timeout)

	topicID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, topicID).Return(nil)

	err := usecase.Delete(context.Background(), topicID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTopicUsecase_Fetch(t *testing.T) {
	mockRepo := new(mocks.TopicRepository)
	timeout := time.Duration(10)
	usecase := test.NewTopicUsecase(mockRepo, timeout)

	fetchedTopics := []domain.Topic{
		{
			ID:                        primitive.NewObjectID(),
			SystemName:                "about-us",
			IncludeInSitemap:          true,
			IncludeInTopMenu:          true,
			IncludeInFooterColumn1:    false,
			IncludeInFooterColumn2:    true,
			IncludeInFooterColumn3:    false,
			DisplayOrder:              1,
			AccessibleWhenStoreClosed: true,
			IsPasswordProtected:       false,
			Password:                  "",
			Title:                     "About Us",
			Body:                      "This is the About Us page content.",
			Published:                 true,
			TopicTemplateID:           primitive.NewObjectID(),
			MetaKeywords:              "about, company, info",
			MetaDescription:           "Learn more about our company.",
			MetaTitle:                 "About Us - Company Info",
			SubjectToAcl:              false,
			LimitedToStores:           false,
			AvailableStartDateTimeUtc: nil,
			AvailableEndDateTimeUtc:   nil,
		},
		{
			ID:                        primitive.NewObjectID(),
			SystemName:                "contact-us",
			IncludeInSitemap:          true,
			IncludeInTopMenu:          false,
			IncludeInFooterColumn1:    true,
			IncludeInFooterColumn2:    false,
			IncludeInFooterColumn3:    true,
			DisplayOrder:              2,
			AccessibleWhenStoreClosed: false,
			IsPasswordProtected:       true,
			Password:                  "securepassword",
			Title:                     "Contact Us",
			Body:                      "This is the Contact Us page content.",
			Published:                 true,
			TopicTemplateID:           primitive.NewObjectID(),
			MetaKeywords:              "contact, support, help",
			MetaDescription:           "Get in touch with us for support.",
			MetaTitle:                 "Contact Us - Support",
			SubjectToAcl:              true,
			LimitedToStores:           true,
			AvailableStartDateTimeUtc: new(time.Time),
			AvailableEndDateTimeUtc:   new(time.Time),
		},
	}
	*fetchedTopics[1].AvailableStartDateTimeUtc = time.Now().AddDate(0, 0, -7) // Available from 7 days ago
	*fetchedTopics[1].AvailableEndDateTimeUtc = time.Now().AddDate(0, 0, 7)    // Available for the next 7 days

	mockRepo.On("Fetch", mock.Anything).Return(fetchedTopics, nil)

	result, err := usecase.Fetch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, fetchedTopics, result)
	mockRepo.AssertExpectations(t)
}
