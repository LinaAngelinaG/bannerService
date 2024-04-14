package bannerRepositoryPostgres

import (
	"bannerService/internal/banner/model"
	"bannerService/pkg/fetch"
	"gitlab.com/piorun102/lg"
)

type BannerRepoInterface interface {
	LoadTokens(ctx lg.CtxLogger) ([]fetch.AccessToken, error)
	GetUserBanner(ctx lg.CtxLogger, req *model.GetUserBannerDB) (*model.BannerDTO, error)
	LoadActiveBanners(ctx lg.CtxLogger, lastLoadedId int64) ([]model.Banner, error)
}
