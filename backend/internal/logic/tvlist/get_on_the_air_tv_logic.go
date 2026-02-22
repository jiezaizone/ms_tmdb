package tvlist

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnTheAirTvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOnTheAirTvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnTheAirTvLogic {
	return &GetOnTheAirTvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOnTheAirTvLogic) GetOnTheAirTv(req *types.PageReq) error {
	// todo: add your logic here and delete this line

	return nil
}
