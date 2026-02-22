package person

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonDetailLogic {
	return &GetPersonDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonDetailLogic) GetPersonDetail(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
