package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieWatchProvidersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieWatchProvidersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieWatchProvidersLogic {
	return &GetMovieWatchProvidersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieWatchProvidersLogic) GetMovieWatchProviders(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
