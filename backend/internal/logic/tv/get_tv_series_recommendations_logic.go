package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesRecommendationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesRecommendationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesRecommendationsLogic {
	return &GetTvSeriesRecommendationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesRecommendationsLogic) GetTvSeriesRecommendations(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
