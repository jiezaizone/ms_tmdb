package discover

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DiscoverMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDiscoverMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DiscoverMovieLogic {
	return &DiscoverMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DiscoverMovieLogic) DiscoverMovie(req *types.DiscoverReq) error {
	// todo: add your logic here and delete this line

	return nil
}
