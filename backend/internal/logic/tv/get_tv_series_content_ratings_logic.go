package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesContentRatingsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesContentRatingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesContentRatingsLogic {
	return &GetTvSeriesContentRatingsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesContentRatingsLogic) GetTvSeriesContentRatings(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
