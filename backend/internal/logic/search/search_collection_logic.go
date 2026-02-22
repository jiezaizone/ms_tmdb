package search

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCollectionLogic {
	return &SearchCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCollectionLogic) SearchCollection(req *types.SearchReq) error {
	// todo: add your logic here and delete this line

	return nil
}
