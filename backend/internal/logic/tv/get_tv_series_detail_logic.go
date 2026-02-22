package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesDetailLogic {
	return &GetTvSeriesDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesDetailLogic) GetTvSeriesDetail(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
