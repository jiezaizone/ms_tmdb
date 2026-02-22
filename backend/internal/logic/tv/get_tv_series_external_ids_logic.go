package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesExternalIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesExternalIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesExternalIdsLogic {
	return &GetTvSeriesExternalIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesExternalIdsLogic) GetTvSeriesExternalIds(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
