package tvseason

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvEpisodeImagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvEpisodeImagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvEpisodeImagesLogic {
	return &GetTvEpisodeImagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvEpisodeImagesLogic) GetTvEpisodeImages(req *types.TvEpisodeReq) error {
	// todo: add your logic here and delete this line

	return nil
}
