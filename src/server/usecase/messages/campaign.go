package usecase

import (
	"context"
	"time"

	domain "earnforglance/server/domain/messages"
)

type campaignUsecase struct {
	campaignRepository domain.CampaignRepository
	contextTimeout     time.Duration
}

func NewCampaignUsecase(campaignRepository domain.CampaignRepository, timeout time.Duration) domain.CampaignUsecase {
	return &campaignUsecase{
		campaignRepository: campaignRepository,
		contextTimeout:     timeout,
	}
}

func (tu *campaignUsecase) Create(c context.Context, campaign *domain.Campaign) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.campaignRepository.Create(ctx, campaign)
}

func (tu *campaignUsecase) Update(c context.Context, campaign *domain.Campaign) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.campaignRepository.Update(ctx, campaign)
}

func (tu *campaignUsecase) Delete(c context.Context, campaign string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.campaignRepository.Delete(ctx, campaign)
}

func (lu *campaignUsecase) FetchByID(c context.Context, campaignID string) (domain.Campaign, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.campaignRepository.FetchByID(ctx, campaignID)
}

func (lu *campaignUsecase) Fetch(c context.Context) ([]domain.Campaign, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.campaignRepository.Fetch(ctx)
}
