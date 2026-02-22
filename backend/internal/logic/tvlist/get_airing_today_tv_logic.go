package tvlist

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAiringTodayTvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAiringTodayTvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAiringTodayTvLogic {
	return &GetAiringTodayTvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAiringTodayTvLogic) GetAiringTodayTv(req *types.PageReq) error {
	// todo: add your logic here and delete this line

	return nil
}
