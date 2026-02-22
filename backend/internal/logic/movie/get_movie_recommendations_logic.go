package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieRecommendationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieRecommendationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieRecommendationsLogic {
	return &GetMovieRecommendationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieRecommendationsLogic) GetMovieRecommendations(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
