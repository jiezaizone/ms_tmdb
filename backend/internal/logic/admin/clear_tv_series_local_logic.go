package admin

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearTvSeriesLocalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearTvSeriesLocalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearTvSeriesLocalLogic {
	return &ClearTvSeriesLocalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClearTvSeriesLocalLogic) ClearTvSeriesLocal(req *types.AdminSyncReq) error {
	// todo: add your logic here and delete this line

	return nil
}
