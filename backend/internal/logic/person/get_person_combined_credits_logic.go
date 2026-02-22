package person

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonCombinedCreditsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonCombinedCreditsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonCombinedCreditsLogic {
	return &GetPersonCombinedCreditsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonCombinedCreditsLogic) GetPersonCombinedCredits(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
