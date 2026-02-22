package misc

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByExternalIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindByExternalIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByExternalIdLogic {
	return &FindByExternalIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindByExternalIdLogic) FindByExternalId(req *types.FindByExternalIdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
