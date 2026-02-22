package movielist

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopRatedMoviesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopRatedMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopRatedMoviesLogic {
	return &GetTopRatedMoviesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopRatedMoviesLogic) GetTopRatedMovies(req *types.PageReq) error {
	// todo: add your logic here and delete this line

	return nil
}
