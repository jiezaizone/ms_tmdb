package movie

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMovieTranslationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieTranslationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieTranslationsLogic {
	return &GetMovieTranslationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieTranslationsLogic) GetMovieTranslations(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
