package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieReleaseDatesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieReleaseDatesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieReleaseDatesLogic {
	return &GetMovieReleaseDatesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieReleaseDatesLogic) GetMovieReleaseDates(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
