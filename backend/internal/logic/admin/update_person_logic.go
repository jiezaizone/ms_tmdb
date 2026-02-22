package admin

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePersonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePersonLogic {
	return &UpdatePersonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePersonLogic) UpdatePerson(req *types.AdminUpdateReq) error {
	// todo: add your logic here and delete this line

	return nil
}
