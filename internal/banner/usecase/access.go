package bannerUsecase

import (
	"bannerService/pkg/fetch"
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"gitlab.com/piorun102/lg"
	"time"
)

func (uc UC) updateAccessTokens() (sleepTime time.Duration) {
	var (
		err error
	)
	sleepTime = accessTokensUPDTime
	sp, c := opentracing.StartSpanFromContext(context.Background(), tokenUPD)
	ext.MessageBusDestination.Set(sp, tokenUPD)
	defer sp.Finish()
	ctx := lg.Ctx(c, nil)
	defer ctx.Send()
	defer ctx.DefError(&err)
	tokens, err := uc.bannerRepo.LoadTokens(ctx)
	if err != nil {
		ctx.Errorf("LoadNewTokens: %v", err)
		return
	}
	fetch.InitAccessTokens(tokens)
	return
}

var (
	lastTokenID = int64(0)
)

const (
	tokenUPD            = "tokenUpdate"
	accessTokensUPDTime = 2 * time.Minute
)
