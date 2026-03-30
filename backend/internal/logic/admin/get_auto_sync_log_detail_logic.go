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

	syncedEntries := convertDetailEntries(detail.Synced)
	syncedPage, syncedPageSize, syncedList := paginateDetailEntries(
		syncedEntries,
		req.SyncedPage,
		req.SyncedPageSize,
	)

	failedEntries := convertDetailEntries(detail.Failed)
	failedPage, failedPageSize, failedList := paginateDetailEntries(
		failedEntries,
		req.FailedPage,
		req.FailedPageSize,
	)

	return &types.AdminAutoSyncLogDetailResp{
		Id:             int64(record.ID),
		TriggeredAt:    formatLogTime(record.TriggeredAt),
		CronExpr:       record.CronExpr,
		Mode:           record.Mode,
		BatchSize:      record.BatchSize,
		Status:         record.Status,
		Checked:        record.Checked,
		Synced:         record.Synced,
		Failed:         record.Failed,
		DurationMs:     record.DurationMs,
		Message:        record.Message,
		StartedAt:      formatLogTime(record.StartedAt),
		FinishedAt:     formatLogTime(record.FinishedAt),
		CreatedAt:      formatLogTime(record.CreatedAt),
		SyncedPage:     syncedPage,
		SyncedPageSize: syncedPageSize,
		SyncedList:     syncedList,
		FailedPage:     failedPage,
		FailedPageSize: failedPageSize,
		FailedList:     failedList,
	}, nil
}

func convertDetailEntries(values []autoSyncExecutionItem) []types.AdminAutoSyncLogDetailEntry {
	if len(values) == 0 {
		return []types.AdminAutoSyncLogDetailEntry{}
	}

	result := make([]types.AdminAutoSyncLogDetailEntry, 0, len(values))
	for _, item := range values {
		result = append(result, types.AdminAutoSyncLogDetailEntry{
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
	return result
}

func paginateDetailEntries(items []types.AdminAutoSyncLogDetailEntry, page, pageSize int) (int, int, []types.AdminAutoSyncLogDetailEntry) {
	page, pageSize = normalizePage(page, pageSize)
	total := len(items)
	if total == 0 {
		return 1, pageSize, []types.AdminAutoSyncLogDetailEntry{}
	}

	totalPages := (total + pageSize - 1) / pageSize
	if page > totalPages {
		page = totalPages
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if end > total {
		end = total
	}

	result := make([]types.AdminAutoSyncLogDetailEntry, end-start)
	copy(result, items[start:end])
	return page, pageSize, result
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
