package admin

import (
	"context"
	"encoding/json"
	"errors"

	"ms_tmdb/internal/model"
	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetAutoSyncLogDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAutoSyncLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAutoSyncLogDetailLogic {
	return &GetAutoSyncLogDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAutoSyncLogDetailLogic) GetAutoSyncLogDetail(req *types.AdminAutoSyncLogDetailReq) (*types.AdminAutoSyncLogDetailResp, error) {
	if req.Id <= 0 {
		return nil, errors.New("日志 ID 不合法")
	}

	var record model.AutoSyncExecutionLog
	if err := l.svcCtx.DB.First(&record, req.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("日志不存在")
		}
		return nil, err
	}

	detail := newAutoSyncExecutionDetail()
	if len(record.Detail) > 0 {
		if err := json.Unmarshal(record.Detail, &detail); err != nil {
			logx.Errorf("解析自动同步日志明细失败: id=%d, err=%v", record.ID, err)
		}
	}

	syncedList := make([]types.AdminAutoSyncLogDetailEntry, 0, len(detail.Synced))
	for _, item := range detail.Synced {
		syncedList = append(syncedList, types.AdminAutoSyncLogDetailEntry{
			MediaType:         item.MediaType,
			TmdbId:            item.TmdbID,
			Name:              item.Name,
			Message:           item.Message,
			RemoteDiffFields:  cloneStringSlice(item.RemoteDiffFields),
			FieldChanges:      convertFieldChanges(item.FieldChanges),
			ChangedFields:     cloneStringSlice(item.ChangedFields),
			OverwrittenFields: cloneStringSlice(item.OverwrittenFields),
			KeptLocalFields:   cloneStringSlice(item.KeptLocalFields),
		})
	}

	failedList := make([]types.AdminAutoSyncLogDetailEntry, 0, len(detail.Failed))
	for _, item := range detail.Failed {
		failedList = append(failedList, types.AdminAutoSyncLogDetailEntry{
			MediaType:         item.MediaType,
			TmdbId:            item.TmdbID,
			Name:              item.Name,
			Message:           item.Message,
			RemoteDiffFields:  cloneStringSlice(item.RemoteDiffFields),
			FieldChanges:      convertFieldChanges(item.FieldChanges),
			ChangedFields:     cloneStringSlice(item.ChangedFields),
			OverwrittenFields: cloneStringSlice(item.OverwrittenFields),
			KeptLocalFields:   cloneStringSlice(item.KeptLocalFields),
		})
	}

	return &types.AdminAutoSyncLogDetailResp{
		Id:          int64(record.ID),
		TriggeredAt: formatLogTime(record.TriggeredAt),
		CronExpr:    record.CronExpr,
		Mode:        record.Mode,
		BatchSize:   record.BatchSize,
		Status:      record.Status,
		Checked:     record.Checked,
		Synced:      record.Synced,
		Failed:      record.Failed,
		DurationMs:  record.DurationMs,
		Message:     record.Message,
		StartedAt:   formatLogTime(record.StartedAt),
		FinishedAt:  formatLogTime(record.FinishedAt),
		CreatedAt:   formatLogTime(record.CreatedAt),
		SyncedList:  syncedList,
		FailedList:  failedList,
	}, nil
}

func cloneStringSlice(values []string) []string {
	if len(values) == 0 {
		return []string{}
	}
	result := make([]string, len(values))
	copy(result, values)
	return result
}

func convertFieldChanges(values []autoSyncFieldChange) []types.AdminAutoSyncLogFieldChange {
	if len(values) == 0 {
		return []types.AdminAutoSyncLogFieldChange{}
	}

	result := make([]types.AdminAutoSyncLogFieldChange, 0, len(values))
	for _, item := range values {
		result = append(result, types.AdminAutoSyncLogFieldChange{
			Field:    item.Field,
			DiffType: item.DiffType,
			Before:   item.Before,
			After:    item.After,
		})
	}
	return result
}
