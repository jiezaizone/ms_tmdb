package discover

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DiscoverTvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDiscoverTvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DiscoverTvLogic {
	return &DiscoverTvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DiscoverTvLogic) DiscoverTv(req *types.DiscoverReq) error {
	// todo: add your logic here and delete this line

	return nil
}
