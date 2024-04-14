package bannerChi

import (
	"bannerService/internal/banner/model"
	"bannerService/pkg/fetch"
	"gitlab.com/piorun102/lg"
)

type (
	I interface {
		GetUserBanner(ctx lg.CtxLogger, req *model.GetUserBannerRequest, role *fetch.Role) (*model.BannerDTO, *fetch.Error)
		GetBannersFiltered(ctx lg.CtxLogger, req *model.GetBannersFilteredRequest, role *fetch.Role) (*model.Banners, *fetch.Error)
		SaveBanner(ctx lg.CtxLogger, req *model.SaveBannerRequest, role *fetch.Role) (*model.Banner, *fetch.Error)
		UpdateBanner(ctx lg.CtxLogger, req *model.UpdateBannerRequest, role *fetch.Role) (*bool, *fetch.Error)
	}
)
