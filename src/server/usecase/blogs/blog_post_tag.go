package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/blogs"
)

type blogposttagUsecase struct {
	blogposttagRepository domain.BlogPostTagRepository
	contextTimeout        time.Duration
}

func NewBlogPostTagUsecase(blogposttagRepository domain.BlogPostTagRepository, timeout time.Duration) domain.BlogPostTagUsecase {
	return &blogposttagUsecase{
		blogposttagRepository: blogposttagRepository,
		contextTimeout:        timeout,
	}
}

func (tu *blogposttagUsecase) Create(c context.Context, blogposttag *domain.BlogPostTag) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogposttagRepository.Create(ctx, blogposttag)
}

func (tu *blogposttagUsecase) Update(c context.Context, blogposttag *domain.BlogPostTag) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogposttagRepository.Update(ctx, blogposttag)
}

func (tu *blogposttagUsecase) Delete(c context.Context, blogposttag *domain.BlogPostTag) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogposttagRepository.Delete(ctx, blogposttag)
}

func (lu *blogposttagUsecase) FetchByID(c context.Context, blogposttagID string) (domain.BlogPostTag, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.blogposttagRepository.FetchByID(ctx, blogposttagID)
}

func (lu *blogposttagUsecase) Fetch(c context.Context) ([]domain.BlogPostTag, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.blogposttagRepository.Fetch(ctx)
}
