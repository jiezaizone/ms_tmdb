package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesTranslationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesTranslationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesTranslationsLogic {
	return &GetTvSeriesTranslationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesTranslationsLogic) GetTvSeriesTranslations(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
