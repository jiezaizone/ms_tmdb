package tvseason

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeasonVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeasonVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeasonVideosLogic {
	return &GetTvSeasonVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeasonVideosLogic) GetTvSeasonVideos(req *types.TvSeasonReq) error {
	// todo: add your logic here and delete this line

	return nil
}
