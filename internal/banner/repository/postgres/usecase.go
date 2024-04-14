package bannerRepositoryPostgres

import (
	"bannerService/internal/banner/model"
	"bannerService/pkg/fetch"
	storageNew "bannerService/pkg/storage"
	"github.com/jackc/pgx/v5"
	"gitlab.com/piorun102/lg"
)

func (br *BannerRepo) LoadActiveBanners(ctx lg.CtxLogger, lastLoadedId int64) ([]model.Banner, error) {
	query := `
      select 
          tag_id, feature_id, banner_content as text, banner_url as url, banner_name as title, active 
      from banner;`
	return storageNew.SelectStructs[model.Banner](ctx, br.db, query, nil)
}

func (br *BannerRepo) GetUserBanner(ctx lg.CtxLogger, req *model.GetUserBannerDB) (*model.BannerDTO, error) {
	query := `
      select 
         banner_content as text, banner_url as url, banner_name as title
      from banner
      where tag_id = @tag and feature_id = @feature and active = @active
      ;`
	args := pgx.NamedArgs{
		"tag":     req.Req.TagId,
		"feature": req.Req.FeatureId,
		"active":  req.Active,
	}
	return storageNew.SelectOneStruct[model.BannerDTO](ctx, br.db, query, args)
}

func (br *BannerRepo) LoadTokens(ctx lg.CtxLogger) ([]fetch.AccessToken, error) {
	query := `
      select 
          access_token as token, name as role, expire_time
      from access join role 
          on access.role_id = role.id
      where active 
        and expire_time > now() at time zone 'Europe/Moscow';`
	return storageNew.SelectStructs[fetch.AccessToken](ctx, br.db, query, nil)
}
