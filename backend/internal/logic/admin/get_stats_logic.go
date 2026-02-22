package admin

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"ms_tmdb/internal/svc"
)

type GetStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatsLogic {
	return &GetStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatsLogic) GetStats() error {
	// todo: add your logic here and delete this line

	return nil
}
