package bannerRepositoryPostgres

import "github.com/jackc/pgx/v5/pgxpool"

type BannerRepo struct {
	db *pgxpool.Pool
}

func NewBannerRepo(db *pgxpool.Pool) BannerRepoInterface {
	return &BannerRepo{
		db: db,
	}
}
