package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesVideosLogic {
	return &GetTvSeriesVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesVideosLogic) GetTvSeriesVideos(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
