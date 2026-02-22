package person

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonMovieCreditsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonMovieCreditsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonMovieCreditsLogic {
	return &GetPersonMovieCreditsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonMovieCreditsLogic) GetPersonMovieCredits(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
