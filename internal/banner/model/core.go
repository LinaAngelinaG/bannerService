package model

import (
	"fmt"
)

type BannerDTO struct {
	Url   string `json:"url"`
	Text  string `json:"text"`
	Title string `json:"title"`
}

type Banner struct {
	Tag     int64  `json:"tag_id" db:"tag_id"`
	Feature int64  `json:"feature_id" db:"feature_id"`
	Url     string `json:"url" db:"url"`
	Content string `json:"text" db:"text"`
	Title   string `json:"title" db:"title"`
	Active  bool   `json:"active" db:"active"`
}

func (b Banner) GetName() string {
	return fmt.Sprintf("%d%d", b.Tag, b.Feature)
}

type Banners struct {
	Banners []Banner
}
