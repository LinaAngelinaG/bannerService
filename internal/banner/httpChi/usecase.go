package bannerChi

import (
	bannerInterface "bannerService/internal/banner/interface"
	"bannerService/internal/banner/model"
	"bannerService/pkg/fetch"
	"github.com/puzpuzpuz/xsync/v3"
	"gitlab.com/piorun102/lg"
	"sync"
)

type UC struct {
	bannerUC bannerInterface.I
	mtx      *xsync.MapOf[string, *sync.Mutex]
}

func (uc UC) GetUserBanner(ctx lg.CtxLogger, req *model.GetUserBannerRequest, role *fetch.Role) (*model.BannerDTO, *fetch.Error) {
	//хранить баннеры в мапе по ключу строке: тэг_фича
	if req.NeedLastVer {
		banner, err := uc.bannerUC.GetUserLastVerBanner(ctx, &model.GetUserBanner{Req: req, Role: role})
		if err != nil {
			return nil, err
		}
		return banner, nil
	}
	banner, err := uc.bannerUC.GetUserBannerFromCache(ctx, &model.GetUserBanner{Req: req, Role: role})
	if err != nil {
		return nil, err
	}
	return banner, nil
}

func (uc UC) GetBannersFiltered(ctx lg.CtxLogger, req *model.GetBannersFilteredRequest, role *fetch.Role) (*model.Banners, *fetch.Error) {
	//TODO implement me
	panic("implement me")
}

func (uc UC) SaveBanner(ctx lg.CtxLogger, req *model.SaveBannerRequest, role *fetch.Role) (*model.Banner, *fetch.Error) {
	//TODO implement me
	panic("implement me")
}

func (uc UC) UpdateBanner(ctx lg.CtxLogger, req *model.UpdateBannerRequest, role *fetch.Role) (*bool, *fetch.Error) {
	//TODO implement me
	panic("implement me")
}

func NewChiUseCase(
	bannerUC bannerInterface.I,
	// TODO delete
	// btcUC usecase.UC,
) I {
	return &UC{
		mtx:      xsync.NewMapOf[string, *sync.Mutex](),
		bannerUC: bannerUC,
	}
}
