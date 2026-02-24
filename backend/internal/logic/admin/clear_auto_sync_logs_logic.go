package admin

import (
	"context"
	"fmt"

	"ms_tmdb/internal/model"
	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ClearAutoSyncLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearAutoSyncLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearAutoSyncLogsLogic {
	return &ClearAutoSyncLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClearAutoSyncLogsLogic) ClearAutoSyncLogs() (*types.AdminAutoSyncLogClearResp, error) {
	tableName, err := autoSyncLogTableName(l.svcCtx.DB)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", tableName)
	if err := l.svcCtx.DB.Exec(sql).Error; err != nil {
		return nil, err
	}

	return &types.AdminAutoSyncLogClearResp{
		Message: "执行日志已清空",
	}, nil
}

func autoSyncLogTableName(db *gorm.DB) (string, error) {
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(&model.AutoSyncExecutionLog{}); err != nil {
		return "", err
	}
	return stmt.Table, nil
}
