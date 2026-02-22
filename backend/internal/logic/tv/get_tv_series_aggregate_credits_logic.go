package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesAggregateCreditsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesAggregateCreditsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesAggregateCreditsLogic {
	return &GetTvSeriesAggregateCreditsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesAggregateCreditsLogic) GetTvSeriesAggregateCredits(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
