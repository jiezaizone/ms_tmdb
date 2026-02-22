package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieImagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieImagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieImagesLogic {
	return &GetMovieImagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieImagesLogic) GetMovieImages(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
