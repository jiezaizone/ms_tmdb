package tvseason

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeasonCreditsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeasonCreditsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeasonCreditsLogic {
	return &GetTvSeasonCreditsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeasonCreditsLogic) GetTvSeasonCredits(req *types.TvSeasonReq) error {
	// todo: add your logic here and delete this line

	return nil
}
