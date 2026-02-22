package misc

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGenreMovieListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGenreMovieListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGenreMovieListLogic {
	return &GetGenreMovieListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGenreMovieListLogic) GetGenreMovieList(req *types.GenreListReq) error {
	// todo: add your logic here and delete this line

	return nil
}
