package bannerInterface

import (
	"bannerService/internal/banner/model"
	"bannerService/pkg/fetch"
	"gitlab.com/piorun102/lg"
)

type I interface {
	GetUserLastVerBanner(ctx lg.CtxLogger, req *model.GetUserBanner) (*model.BannerDTO, *fetch.Error)
	GetUserBannerFromCache(ctx lg.CtxLogger, req *model.GetUserBanner) (*model.BannerDTO, *fetch.Error)
}
