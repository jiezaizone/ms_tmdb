package trending

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTrendingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTrendingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTrendingLogic {
	return &GetTrendingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTrendingLogic) GetTrending(req *types.TrendingReq) error {
	// todo: add your logic here and delete this line

	return nil
}
