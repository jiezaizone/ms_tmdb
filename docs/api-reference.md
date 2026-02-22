# MS-TMDB API 参考文档（当前实现）

更新时间：2026-02-21

## 基本信息

- 服务地址：`http://localhost:8888`
- TMDB 代理前缀：`/api/v3`
- 管理接口前缀：`/api/admin`
- 响应格式：`application/json`

说明：

1. `/api/v3/*` 默认返回 TMDB 原始 JSON。
2. 无需客户端传 `api_key`，服务端会自动附加配置中的 TMDB Key。
3. `language`、`page`、`region`、`append_to_response` 及其他查询参数会按规则透传。

---

## 一、TMDB 代理接口（`/api/v3`）

### 1.1 带本地读穿缓存的详情接口

- `GET /api/v3/movie/{id}`
- `GET /api/v3/tv/{id}`
- `GET /api/v3/person/{id}`

行为：

1. 本地命中且未过期（或已本地修改）时，直接返回本地 `tmdb_data`。
2. 否则回源 TMDB，成功后写入本地并返回。
3. TMDB 不可用且本地有缓存时，降级返回本地缓存。

### 1.2 透传接口

除上述 3 个详情接口外，以下路径由服务直接转发到 TMDB：

- 电影：
  - `/api/v3/movie/{id}/credits`
  - `/api/v3/movie/{id}/images`
  - `/api/v3/movie/{id}/videos`
  - `/api/v3/movie/{id}/keywords`
  - `/api/v3/movie/{id}/similar`
  - `/api/v3/movie/{id}/recommendations`
  - `/api/v3/movie/{id}/external_ids`
  - `/api/v3/movie/{id}/translations`
  - `/api/v3/movie/{id}/release_dates`
  - `/api/v3/movie/{id}/watch/providers`
  - `/api/v3/movie/{id}/alternative_titles`
  - `/api/v3/movie/now_playing`
  - `/api/v3/movie/popular`
  - `/api/v3/movie/top_rated`
  - `/api/v3/movie/upcoming`
- 电视剧：
  - `/api/v3/tv/{id}/credits`
  - `/api/v3/tv/{id}/aggregate_credits`
  - `/api/v3/tv/{id}/images`
  - `/api/v3/tv/{id}/videos`
  - `/api/v3/tv/{id}/keywords`
  - `/api/v3/tv/{id}/similar`
  - `/api/v3/tv/{id}/recommendations`
  - `/api/v3/tv/{id}/external_ids`
  - `/api/v3/tv/{id}/content_ratings`
  - `/api/v3/tv/{id}/translations`
  - `/api/v3/tv/{id}/watch/providers`
  - `/api/v3/tv/{series_id}/season/{season_number}`
  - `/api/v3/tv/{series_id}/season/{season_number}/credits`
  - `/api/v3/tv/{series_id}/season/{season_number}/images`
  - `/api/v3/tv/{series_id}/season/{season_number}/videos`
  - `/api/v3/tv/{series_id}/season/{season_number}/episode/{episode_number}`
  - `/api/v3/tv/{series_id}/season/{season_number}/episode/{episode_number}/credits`
  - `/api/v3/tv/{series_id}/season/{season_number}/episode/{episode_number}/images`
  - `/api/v3/tv/airing_today`
  - `/api/v3/tv/on_the_air`
  - `/api/v3/tv/popular`
  - `/api/v3/tv/top_rated`
- 人物：
  - `/api/v3/person/{id}/movie_credits`
  - `/api/v3/person/{id}/tv_credits`
  - `/api/v3/person/{id}/combined_credits`
  - `/api/v3/person/{id}/images`
  - `/api/v3/person/{id}/external_ids`
  - `/api/v3/person/popular`
- 搜索：
  - `/api/v3/search/movie`
  - `/api/v3/search/tv`
  - `/api/v3/search/person`
  - `/api/v3/search/multi`
  - `/api/v3/search/keyword`
  - `/api/v3/search/collection`
  - `/api/v3/search/company`
- 发现/趋势/其他：
  - `/api/v3/discover/movie`
  - `/api/v3/discover/tv`
  - `/api/v3/trending/{media_type}/{time_window}`
  - `/api/v3/genre/movie/list`
  - `/api/v3/genre/tv/list`
  - `/api/v3/configuration`
  - `/api/v3/find/{external_id}`（需 `external_source`）

提示：未明确列出的 `/api/v3/*` 路径也会尝试透传 TMDB。

---

## 二、管理接口（`/api/admin`）

### 2.1 修改本地覆盖数据

- `PUT /api/admin/movie/{id}`
- `PUT /api/admin/tv/{id}`
- `PUT /api/admin/person/{id}`

请求体：任意 JSON 对象，写入 `local_data`，并将 `is_modified=true`。
返回值：`tmdb_data` 与 `local_data` 的浅合并结果（同名字段以 `local_data` 覆盖）。

### 2.2 强制同步上游数据

- `POST /api/admin/sync/movie/{id}`
- `POST /api/admin/sync/tv/{id}`
- `POST /api/admin/sync/person/{id}`

行为：回源 TMDB，覆盖本地 `tmdb_data`，清空 `local_data`，重置 `is_modified=false`。

### 2.3 清理本地覆盖

- `DELETE /api/admin/movie/{id}/local`
- `DELETE /api/admin/tv/{id}/local`

行为：清空 `local_data`，并重置 `is_modified=false`。

### 2.4 统计

- `GET /api/admin/stats`

示例响应：

```json
{
  "movies": 100,
  "tv_series": 80,
  "people": 200,
  "movies_modified": 5,
  "tv_series_modified": 2
}
```

---

## 三、错误语义

- TMDB 代理失败：返回 `502`，结构为

```json
{
  "success": false,
  "status_code": 502,
  "status_message": "具体错误信息"
}
```

- 管理接口业务错误（如资源不存在）：返回 `400`，结构为

```json
{
  "error": "错误信息"
}
```
