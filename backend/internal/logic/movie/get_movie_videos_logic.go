package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieVideosLogic {
	return &GetMovieVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieVideosLogic) GetMovieVideos(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
