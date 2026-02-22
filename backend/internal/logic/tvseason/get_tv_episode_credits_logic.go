package tvseason

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvEpisodeCreditsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvEpisodeCreditsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvEpisodeCreditsLogic {
	return &GetTvEpisodeCreditsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvEpisodeCreditsLogic) GetTvEpisodeCredits(req *types.TvEpisodeReq) error {
	// todo: add your logic here and delete this line

	return nil
}
