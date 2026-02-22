package person

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonExternalIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonExternalIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonExternalIdsLogic {
	return &GetPersonExternalIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonExternalIdsLogic) GetPersonExternalIds(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
