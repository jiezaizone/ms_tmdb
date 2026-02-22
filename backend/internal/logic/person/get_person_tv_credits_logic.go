package person

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonTvCreditsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonTvCreditsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonTvCreditsLogic {
	return &GetPersonTvCreditsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonTvCreditsLogic) GetPersonTvCredits(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
