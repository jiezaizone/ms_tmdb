package misc

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"ms_tmdb/internal/svc"
)

type GetConfigurationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationLogic {
	return &GetConfigurationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigurationLogic) GetConfiguration() error {
	// todo: add your logic here and delete this line

	return nil
}
