package search

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchPersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPersonLogic {
	return &SearchPersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchPersonLogic) SearchPerson(req *types.SearchReq) error {
	// todo: add your logic here and delete this line

	return nil
}
