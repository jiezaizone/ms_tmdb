package admin

import (
	"encoding/json"
	"fmt"
	"strings"

	"ms_tmdb/internal/model"

	"gorm.io/gorm"
)

// normalizePage 规范化分页参数
func normalizePage(page, pageSize int) (int, int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	return page, pageSize
}

// normalizeSearchMode 规范化搜索模式，默认 contains
func normalizeSearchMode(mode string) string {
	switch strings.ToLower(strings.TrimSpace(mode)) {
	case "prefix":
		return "prefix"
	default:
		return "contains"
	}
}

// applyKeywordFilter 在给定字段上应用模糊匹配
func applyKeywordFilter(db *gorm.DB, keyword, mode string, columns ...string) *gorm.DB {
	kw := strings.TrimSpace(keyword)
	if kw == "" || len(columns) == 0 {
		return db
	}

	pattern := "%" + kw + "%"
	if normalizeSearchMode(mode) == "prefix" {
		pattern = kw + "%"
	}

	parts := make([]string, 0, len(columns))
	args := make([]interface{}, 0, len(columns))
	for _, col := range columns {
		parts = append(parts, fmt.Sprintf("%s ILIKE ?", col))
		args = append(args, pattern)
	}

	return db.Where(strings.Join(parts, " OR "), args...)
}

// mergeTMDBData 把 patch 合并到原始 TMDB 数据中
func mergeTMDBData(base model.RawJSON, patch map[string]interface{}) (model.RawJSON, error) {
	if len(patch) == 0 {
		return base, nil
	}

	payload := make(map[string]interface{})
	if len(base) > 0 {
		if err := json.Unmarshal(base, &payload); err != nil {
			return nil, err
		}
	}

	for key, value := range patch {
		payload[key] = value
	}

	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return model.RawJSON(raw), nil
}

// toRawJSON 把任意对象转为 RawJSON
func toRawJSON(v interface{}) (model.RawJSON, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return model.RawJSON(raw), nil
}

// trimPtrString 去除字符串指针值的前后空格
func trimPtrString(v *string) string {
	if v == nil {
		return ""
	}
	return strings.TrimSpace(*v)
}

// buildGenresFromNames 把类型名转换为详情页可渲染结构
func buildGenresFromNames(names []string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(names))
	seen := make(map[string]struct{}, len(names))
	id := 1
	for _, raw := range names {
		name := strings.TrimSpace(raw)
		if name == "" {
			continue
		}
		key := strings.ToLower(name)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, map[string]interface{}{
			"id":   id,
			"name": name,
		})
		id++
	}
	return result
}
