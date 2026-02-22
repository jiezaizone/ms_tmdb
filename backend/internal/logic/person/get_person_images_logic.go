package person

import (
	"context"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonImagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonImagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonImagesLogic {
	return &GetPersonImagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonImagesLogic) GetPersonImages(req *types.DetailReq) error {
	// todo: add your logic here and delete this line

	return nil
}
