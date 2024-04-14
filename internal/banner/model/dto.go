package model

import (
	"bannerService/pkg/fetch"
)

type GetUserBannerRequest struct {
	TagId       int64 `json:"tag_id" required`
	FeatureId   int64 `json:"feature_id" required`
	NeedLastVer bool  `json:"use_last_revision" omitempty`
}

type GetUserBanner struct {
	Req  *GetUserBannerRequest
	Role *fetch.Role
}

type GetUserBannerDB struct {
	Req    *GetUserBannerRequest
	Active bool
}

type GetBannersFilteredRequest struct {
}

type UpdateBannerRequest struct {
}

type SaveBannerRequest struct {
}
