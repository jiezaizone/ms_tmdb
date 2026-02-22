package misc

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGenreTvListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGenreTvListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGenreTvListLogic {
	return &GetGenreTvListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGenreTvListLogic) GetGenreTvList(req *types.GenreListReq) error {
	// todo: add your logic here and delete this line

	return nil
}
