package admin

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncMovieLogic {
	return &SyncMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncMovieLogic) SyncMovie(req *types.AdminSyncReq) error {
	// todo: add your logic here and delete this line

	return nil
}
