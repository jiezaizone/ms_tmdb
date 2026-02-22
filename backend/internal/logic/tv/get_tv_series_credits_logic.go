package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesCreditsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesCreditsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesCreditsLogic {
	return &GetTvSeriesCreditsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesCreditsLogic) GetTvSeriesCredits(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
