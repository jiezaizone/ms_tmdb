package admin

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncTvSeriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncTvSeriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncTvSeriesLogic {
	return &SyncTvSeriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncTvSeriesLogic) SyncTvSeries(req *types.AdminSyncReq) error {
	// todo: add your logic here and delete this line

	return nil
}
