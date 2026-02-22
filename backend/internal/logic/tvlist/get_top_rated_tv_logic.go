package tvlist

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopRatedTvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopRatedTvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopRatedTvLogic {
	return &GetTopRatedTvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopRatedTvLogic) GetTopRatedTv(req *types.PageReq) error {
	// todo: add your logic here and delete this line

	return nil
}
