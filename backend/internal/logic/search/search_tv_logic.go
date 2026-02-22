package search

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchTvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTvLogic {
	return &SearchTvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchTvLogic) SearchTv(req *types.SearchReq) error {
	// todo: add your logic here and delete this line

	return nil
}
