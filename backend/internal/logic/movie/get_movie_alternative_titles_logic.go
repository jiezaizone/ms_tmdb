package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieAlternativeTitlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieAlternativeTitlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieAlternativeTitlesLogic {
	return &GetMovieAlternativeTitlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieAlternativeTitlesLogic) GetMovieAlternativeTitles(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
