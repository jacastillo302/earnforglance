package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/blogs"
)

type blogcommentUsecase struct {
	blogcommentRepository domain.BlogCommentRepository
	contextTimeout        time.Duration
}

func NewBlogCommentUsecase(blogcommentRepository domain.BlogCommentRepository, timeout time.Duration) domain.BlogCommentUsecase {
	return &blogcommentUsecase{
		blogcommentRepository: blogcommentRepository,
		contextTimeout:        timeout,
	}
}

func (tu *blogcommentUsecase) Create(c context.Context, blogComment *domain.BlogComment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogcommentRepository.Create(ctx, blogComment)
}

func (tu *blogcommentUsecase) Update(c context.Context, blogComment *domain.BlogComment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogcommentRepository.Update(ctx, blogComment)
}

func (tu *blogcommentUsecase) Delete(c context.Context, blogComment *domain.BlogComment) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.blogcommentRepository.Delete(ctx, blogComment)
}

func (lu *blogcommentUsecase) FetchByID(c context.Context, blogCommentID string) (domain.BlogComment, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.blogcommentRepository.FetchByID(ctx, blogCommentID)
}

func (lu *blogcommentUsecase) Fetch(c context.Context) ([]domain.BlogComment, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.blogcommentRepository.Fetch(ctx)
}
