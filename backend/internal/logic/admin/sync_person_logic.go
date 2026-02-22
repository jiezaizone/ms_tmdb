package admin

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncPersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncPersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncPersonLogic {
	return &SyncPersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncPersonLogic) SyncPerson(req *types.AdminSyncReq) error {
	// todo: add your logic here and delete this line

	return nil
}
