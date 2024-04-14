package bannerUsecase

import (
	"bannerService/internal/banner/model"
	"bannerService/pkg/fetch"
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"gitlab.com/piorun102/lg"
	"net/http"
	"strings"
	"time"
)

const (
	bannerLoading  = "bannerLoading"
	bannersUPDTime = 10 * time.Minute
	bannerNotFound = "Баннер для не найден"
	internalError  = "Внутренняя ошибка сервера"
)

var (
	internalErr = fetch.Error{Message: internalError, Code: http.StatusInternalServerError}
	notFoundErr = fetch.Error{Message: bannerNotFound, Code: http.StatusNotFound}
)

func (uc *UC) GetUserLastVerBanner(ctx lg.CtxLogger, req *model.GetUserBanner) (*model.BannerDTO, *fetch.Error) {
	active := false
	if *req.Role == fetch.Admin {
		active = true
	}
	banner, err := uc.bannerRepo.GetUserBanner(ctx, &model.GetUserBannerDB{Req: req.Req, Active: active})
	if err != nil {
		lg.Errorf("GetUserBanner failed: %+v", err)
		if !strings.Contains(err.Error(), "no rows in result set") {
			return nil, &notFoundErr
		}
		return nil, &internalErr
	}
	return banner, nil
}

func (uc *UC) GetUserBannerFromCache(ctx lg.CtxLogger, req *model.GetUserBanner) (*model.BannerDTO, *fetch.Error) {
	b := model.Banner{Tag: req.Req.TagId, Feature: req.Req.FeatureId}
	banner, ok := uc.activeBanners.Load(b.GetName())
	if ok {
		return &model.BannerDTO{Title: banner.Title, Url: banner.Url, Text: banner.Content}, nil
	}
	return nil, &notFoundErr
}

func (uc *UC) updateBanners() (sleepTime time.Duration) {
	var (
		err error
	)
	sleepTime = bannersUPDTime
	sp, c := opentracing.StartSpanFromContext(context.Background(), bannerLoading)
	ext.MessageBusDestination.Set(sp, bannerLoading)
	defer sp.Finish()
	ctx := lg.Ctx(c, nil)
	defer ctx.Send()
	defer ctx.DefError(&err)
	banners, err := uc.bannerRepo.LoadActiveBanners(ctx, lastTokenID)
	if err != nil {
		ctx.Errorf("LoadNewTokens: %v", err)
		return
	}
	for _, banner := range banners {
		uc.activeBanners.Store(banner.GetName(), banner)
	}
	return
}
