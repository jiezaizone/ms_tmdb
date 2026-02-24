package admin

import (
	"strings"

	"ms_tmdb/internal/types"
)

type autoSyncExecutionDetail struct {
	Synced []autoSyncExecutionItem `json:"synced"`
	Failed []autoSyncExecutionItem `json:"failed"`
}

type autoSyncExecutionItem struct {
	MediaType         string                `json:"media_type"`
	TmdbID            int                   `json:"tmdb_id"`
	Name              string                `json:"name"`
	Message           string                `json:"message"`
	RemoteDiffFields  []string              `json:"remote_diff_fields,omitempty"`
	FieldChanges      []autoSyncFieldChange `json:"field_changes,omitempty"`
	ChangedFields     []string              `json:"changed_fields,omitempty"`
	OverwrittenFields []string              `json:"overwritten_fields,omitempty"`
	KeptLocalFields   []string              `json:"kept_local_fields,omitempty"`
}

type autoSyncFieldChange struct {
	Field    string `json:"field"`
	DiffType string `json:"diff_type"`
	Before   string `json:"before"`
	After    string `json:"after"`
}

func newAutoSyncExecutionDetail() autoSyncExecutionDetail {
	return autoSyncExecutionDetail{
		Synced: make([]autoSyncExecutionItem, 0),
		Failed: make([]autoSyncExecutionItem, 0),
	}
}

func convertCompareFieldChanges(details []types.AdminCompareFieldDetail) []autoSyncFieldChange {
	if len(details) == 0 {
		return []autoSyncFieldChange{}
	}

	changes := make([]autoSyncFieldChange, 0, len(details))
	for _, item := range details {
		if strings.TrimSpace(item.Local) == strings.TrimSpace(item.Remote) {
			continue
		}

		changes = append(changes, autoSyncFieldChange{
			Field:    item.Field,
			DiffType: item.DiffType,
			Before:   item.Local,
			After:    item.Remote,
		})
	}
	return changes
}
