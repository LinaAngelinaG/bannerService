package server

import (
	"bannerService/config/inj"
	httpChi "bannerService/internal/banner/httpChi"
	bannerRepository "bannerService/internal/banner/repository/postgres"
	bannerUsecase "bannerService/internal/banner/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-resty/resty/v2"
	"time"
)

const (
	oneMinuteTime                 = time.Minute
	comparisonSendDurationMinutes = 10 * time.Minute
	daySendMessageTime            = 12
	eveningSendMessageTime        = 19
	discrepancyAlertTime          = 3 * time.Minute
	oneDayTime                    = 24 * time.Hour
	startTransactions             = 1
)

func (s *Server) MapHandlers() error {
	var (
		r  = chi.NewRouter()
		rs *resty.Client
	)
	defer httpChi.HandlePanic()
	bannerRepo := bannerRepository.NewBannerRepo(inj.I.BannerDB())
	coreUC := bannerUsecase.NewCoreUseCase(bannerRepo, rs)

	r.Use(middleware.Logger)
	r.Use(Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	}))

	chiUC := httpChi.NewChiUseCase(coreUC)
	httpChi.New(r, chiUC).InitRoutes(s.ctx)

	s.httpServer.Handler = r

	return nil
}
