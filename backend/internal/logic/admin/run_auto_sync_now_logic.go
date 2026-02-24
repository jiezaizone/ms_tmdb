package admin

import (
	"context"
	"errors"

	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RunAutoSyncNowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRunAutoSyncNowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RunAutoSyncNowLogic {
	return &RunAutoSyncNowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RunAutoSyncNowLogic) RunAutoSyncNow() (*types.AdminAutoSyncRunResp, error) {
	scheduler := GetLibraryAutoSyncScheduler()
	if scheduler == nil {
		return nil, errors.New("自动同步调度器未初始化")
	}

	started, running, err := scheduler.TriggerNow()
	if err != nil {
		return nil, err
	}
	if !started {
		return &types.AdminAutoSyncRunResp{
			Started: false,
			Running: running,
			Message: "任务正在执行中，请稍后再试",
		}, nil
	}

	return &types.AdminAutoSyncRunResp{
		Started: true,
		Running: running,
		Message: "已触发一次立即同步任务，执行完成后可在日志中查看结果",
	}, nil
}
