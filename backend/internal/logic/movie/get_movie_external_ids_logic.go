package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieExternalIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieExternalIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieExternalIdsLogic {
	return &GetMovieExternalIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieExternalIdsLogic) GetMovieExternalIds(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
