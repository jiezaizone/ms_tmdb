package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieSimilarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieSimilarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieSimilarLogic {
	return &GetMovieSimilarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieSimilarLogic) GetMovieSimilar(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
