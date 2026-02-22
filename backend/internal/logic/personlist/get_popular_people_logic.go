package personlist

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPopularPeopleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPopularPeopleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPopularPeopleLogic {
	return &GetPopularPeopleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPopularPeopleLogic) GetPopularPeople(req *types.PageReq) error {
	// todo: add your logic here and delete this line

	return nil
}
