package tvseason

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTvSeasonDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTvSeasonDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTvSeasonDetailLogic {
	return &GetTvSeasonDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTvSeasonDetailLogic) GetTvSeasonDetail(req *types.TvSeasonReq) error {
	// todo: add your logic here and delete this line

	return nil
}
