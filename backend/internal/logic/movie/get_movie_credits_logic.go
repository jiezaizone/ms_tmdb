package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieCreditsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieCreditsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieCreditsLogic {
	return &GetMovieCreditsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieCreditsLogic) GetMovieCredits(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
