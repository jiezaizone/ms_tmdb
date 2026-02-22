package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesWatchProvidersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesWatchProvidersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesWatchProvidersLogic {
	return &GetTvSeriesWatchProvidersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesWatchProvidersLogic) GetTvSeriesWatchProviders(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
