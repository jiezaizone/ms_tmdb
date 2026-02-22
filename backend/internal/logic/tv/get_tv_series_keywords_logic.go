package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesKeywordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesKeywordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesKeywordsLogic {
	return &GetTvSeriesKeywordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesKeywordsLogic) GetTvSeriesKeywords(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
