package search

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchKeywordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchKeywordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchKeywordLogic {
	return &SearchKeywordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchKeywordLogic) SearchKeyword(req *types.SearchReq) error {
	// todo: add your logic here and delete this line

	return nil
}
