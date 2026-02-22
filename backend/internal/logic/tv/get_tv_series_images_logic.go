package tv

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeriesImagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeriesImagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeriesImagesLogic {
	return &GetTvSeriesImagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeriesImagesLogic) GetTvSeriesImages(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
