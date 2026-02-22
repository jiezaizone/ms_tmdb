package admin

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearMovieLocalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearMovieLocalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearMovieLocalLogic {
	return &ClearMovieLocalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClearMovieLocalLogic) ClearMovieLocal(req *types.AdminSyncReq) error {
	// todo: add your logic here and delete this line

	return nil
}
