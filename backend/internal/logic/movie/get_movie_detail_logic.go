package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieDetailLogic {
	return &GetMovieDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieDetailLogic) GetMovieDetail(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
