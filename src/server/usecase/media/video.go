package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/media"
)

type videoUsecase struct {
	videoRepository domain.VideoRepository
	contextTimeout  time.Duration
}

func NewVideoUsecase(videoRepository domain.VideoRepository, timeout time.Duration) domain.VideoUsecase {
	return &videoUsecase{
		videoRepository: videoRepository,
		contextTimeout:  timeout,
	}
}

func (vu *videoUsecase) Create(c context.Context, video *domain.Video) error {
	ctx, cancel := context.WithTimeout(c, vu.contextTimeout)
	defer cancel()
	return vu.videoRepository.Create(ctx, video)
}

func (vu *videoUsecase) Update(c context.Context, video *domain.Video) error {
	ctx, cancel := context.WithTimeout(c, vu.contextTimeout)
	defer cancel()
	return vu.videoRepository.Update(ctx, video)
}

func (vu *videoUsecase) Delete(c context.Context, video string) error {
	ctx, cancel := context.WithTimeout(c, vu.contextTimeout)
	defer cancel()
	return vu.videoRepository.Delete(ctx, video)
}

func (vu *videoUsecase) FetchByID(c context.Context, videoID string) (domain.Video, error) {
	ctx, cancel := context.WithTimeout(c, vu.contextTimeout)
	defer cancel()
	return vu.videoRepository.FetchByID(ctx, videoID)
}

func (vu *videoUsecase) Fetch(c context.Context) ([]domain.Video, error) {
	ctx, cancel := context.WithTimeout(c, vu.contextTimeout)
	defer cancel()
	return vu.videoRepository.Fetch(ctx)
}
