package movielist

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPopularMoviesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPopularMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPopularMoviesLogic {
	return &GetPopularMoviesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPopularMoviesLogic) GetPopularMovies(req *types.PageReq) error {
	// todo: add your logic here and delete this line

	return nil
}
