package movielist

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpcomingMoviesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUpcomingMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpcomingMoviesLogic {
	return &GetUpcomingMoviesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUpcomingMoviesLogic) GetUpcomingMovies(req *types.PageReq) error {
	// todo: add your logic here and delete this line

	return nil
}
