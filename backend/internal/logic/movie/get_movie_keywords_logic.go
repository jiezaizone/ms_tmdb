package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieKeywordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieKeywordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieKeywordsLogic {
	return &GetMovieKeywordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieKeywordsLogic) GetMovieKeywords(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
