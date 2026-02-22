package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesSimilarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesSimilarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesSimilarLogic {
	return &GetTvSeriesSimilarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesSimilarLogic) GetTvSeriesSimilar(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
