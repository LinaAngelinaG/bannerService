package bannerUsecase

import (
	bannerInterface "bannerService/internal/banner/interface"
	"bannerService/internal/banner/model"
	bannerRepositoryPostgres "bannerService/internal/banner/repository/postgres"
	"bannerService/internal/server/schedule"
	"github.com/go-resty/resty/v2"
	"github.com/puzpuzpuz/xsync/v3"
)

type UC struct {
	rs *resty.Client
	//TODO update logic for types of map to use it fast
	activeBanners *xsync.MapOf[string, model.Banner]
	bannerRepo    bannerRepositoryPostgres.BannerRepoInterface
}

func NewCoreUseCase(
	bannerRepo bannerRepositoryPostgres.BannerRepoInterface,
	restyC *resty.Client,
) bannerInterface.I {
	uc := UC{
		rs:            restyC,
		bannerRepo:    bannerRepo,
		activeBanners: xsync.NewMapOf[string, model.Banner](),
	}

	//инициализация аксесс токенов
	schedule.Repeat(uc.updateAccessTokens)
	//инициализация баннеров
	//обновление баннеров из базы
	schedule.Repeat(uc.updateBanners)

	return &uc
}
