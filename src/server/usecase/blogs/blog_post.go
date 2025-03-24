package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/blogs"
)

type blogpostUsecase struct {
	blogpostRepository domain.BlogPostRepository
	contextTimeout     time.Duration
}

func NewBlogPostUsecase(blogpostRepository domain.BlogPostRepository, timeout time.Duration) domain.BlogPostUsecase {
	return &blogpostUsecase{
		blogpostRepository: blogpostRepository,
		contextTimeout:     timeout,
	}
}

func (tu *blogpostUsecase) CreateMany(c context.Context, items []domain.BlogPost) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogpostRepository.CreateMany(ctx, items)
}

func (tu *blogpostUsecase) Create(c context.Context, blogpost *domain.BlogPost) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogpostRepository.Create(ctx, blogpost)
}

func (tu *blogpostUsecase) Update(c context.Context, blogpost *domain.BlogPost) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogpostRepository.Update(ctx, blogpost)
}

func (tu *blogpostUsecase) Delete(c context.Context, blogpost string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogpostRepository.Delete(ctx, blogpost)
}

func (lu *blogpostUsecase) FetchByID(c context.Context, blogpostID string) (domain.BlogPost, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.blogpostRepository.FetchByID(ctx, blogpostID)
}

func (lu *blogpostUsecase) Fetch(c context.Context) ([]domain.BlogPost, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.blogpostRepository.Fetch(ctx)
}
