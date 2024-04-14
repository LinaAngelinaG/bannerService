package bannerChi

import (
	"bannerService/internal/banner/model"
	"bannerService/pkg/fetch"
	"context"
	"gitlab.com/piorun102/lg"
	//lg "example.com/logger"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	router *chi.Mux
	uc     I
}

func New(router *chi.Mux, uc I) *Handler {
	return &Handler{
		router: router,
		uc:     uc,
	}
}

func (h *Handler) InitRoutes(ctx context.Context) {
	fetch.Get[model.GetUserBannerRequest, model.BannerDTO]("/user_banner", h.router, fetch.UserAccess, h.uc.GetUserBanner)
	fetch.Get[model.GetBannersFilteredRequest, model.Banners]("/banner", h.router, fetch.AdminAccess, h.uc.GetBannersFiltered)

	fetch.Post[model.SaveBannerRequest, model.Banner]("/banner", h.router, fetch.AdminAccess, h.uc.SaveBanner)
	fetch.Patch[model.UpdateBannerRequest, bool]("/banner/{id}", h.router, fetch.AdminAccess, h.uc.UpdateBanner)
}

func HandlePanic() {
	if r := recover(); r != nil {
		lg.Ctx(context.Background(), nil).Panicf("Panic occured:", r)
	}
}
