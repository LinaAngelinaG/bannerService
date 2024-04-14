package inj

import (
	"bannerService/config"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/piorun102/lg"
	//lg "example.com/logger"
	"gitlab.com/piorun102/pkg/storage"
)

var I Inj

type (
	inj struct {
		bannerDB *pgxpool.Pool
	}
	Inj interface {
		BannerDB() *pgxpool.Pool
	}
)

func (i *inj) BannerDB() *pgxpool.Pool {
	return i.bannerDB
}

func Init(cfg *config.Config) {
	var (
		i   = &inj{}
		err error
	)

	lg.Connect(cfg.Shared.Logger.Address)
	lg.Tracef("Logger connected")

	lg.Tracef("Nats connected")

	i.bannerDB, err = storage.NewDB(context.TODO(), &storage.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "3846936720",
		Db:       "BitcoinTransactions",
	})
	if err != nil {
		lg.Fatal(err.Error())
	}

	lg.Tracef("Postgres connected")

	I = i

	lg.Debug("All components connected successfully")
}
