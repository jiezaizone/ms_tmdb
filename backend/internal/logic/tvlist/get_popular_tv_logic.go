package tvlist

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPopularTvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPopularTvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPopularTvLogic {
	return &GetPopularTvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPopularTvLogic) GetPopularTv(req *types.PageReq) error {
	// todo: add your logic here and delete this line

	return nil
}
