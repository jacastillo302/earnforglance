package domain

import (
	"context"
	domain "earnforglance/server/domain/topics"
)

type TopicRequest struct {
	ID               string
	Filters          []Filter
	Sort             string
	Limit            int
	Page             int
	Lang             string
	IncludeInTopMenu bool
	Password         string
	Content          []string
}

type TopicResponse struct {
	Topic    domain.Topic
	Template domain.TopicTemplate
}

type TopicsResponse struct {
	Topics []TopicResponse
}

type TopicRepository interface {
	GetTopics(c context.Context, filter TopicRequest) ([]TopicsResponse, error)
}

type TopictUsecase interface {
	GetTopics(c context.Context, filter TopicRequest) ([]TopicsResponse, error)
}
