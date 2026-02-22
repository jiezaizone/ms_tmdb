package tvseason

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvEpisodeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvEpisodeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvEpisodeDetailLogic {
	return &GetTvEpisodeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvEpisodeDetailLogic) GetTvEpisodeDetail(req *types.TvEpisodeReq) error {
	// todo: add your logic here and delete this line

	return nil
}
