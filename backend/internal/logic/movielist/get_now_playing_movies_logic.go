package movielist

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNowPlayingMoviesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNowPlayingMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNowPlayingMoviesLogic {
	return &GetNowPlayingMoviesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNowPlayingMoviesLogic) GetNowPlayingMovies(req *types.PageReq) error {
	// todo: add your logic here and delete this line

	return nil
}
