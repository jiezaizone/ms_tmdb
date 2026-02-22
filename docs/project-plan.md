# MS-TMDB 项目方案（当前实施版）

更新时间：2026-02-21

## 一、项目定位

`MS-TMDB` 是一个基于 `go-zero + PostgreSQL` 的 TMDB 代理服务。
当前版本采用以下固定方案：

1. 对外暴露 `/api/v3/*`，由中间件统一代理到 TMDB。
2. 对电影/电视剧/人物详情接口提供本地 Read-Through 缓存（PostgreSQL）。
3. 通过 `/api/admin/*` 提供本地覆盖、同步、清理、统计能力。

说明：当前不切换实现路线，继续在该方案上迭代。

---

## 二、技术栈

| 组件 | 技术 | 用途 |
|------|------|------|
| API 框架 | go-zero | HTTP 服务与路由注册 |
| 数据库 | PostgreSQL | 存储详情缓存与本地覆盖数据 |
| ORM | GORM | 模型管理与自动迁移 |
| 上游客户端 | 自研 `backend/pkg/tmdbclient` | 请求 TMDB v3 API |
| 配置 | YAML | 服务与连接配置 |

---

## 三、当前架构

```text
客户端
  -> go-zero HTTP Server
    -> /api/v3/*: TmdbProxyMiddleware
      -> ProxyService（详情读穿缓存）
      -> TmdbClient（其余接口透传）
    -> /api/admin/*: Admin Handler
      -> PostgreSQL（local_data / tmdb_data 管理）
```

### 3.1 请求处理策略

- `/api/v3/movie/{id}`、`/api/v3/tv/{id}`、`/api/v3/person/{id}`：
  先查本地，未命中或过期时回源 TMDB，再写回本地。
- 其他 `/api/v3/*`：
  直接透传到 TMDB，返回原始 JSON。
- `/api/admin/*`：
  管理本地覆盖数据与同步行为。

### 3.2 数据字段策略

每个核心模型包含：

- `tmdb_data`：TMDB 原始响应（JSONB）
- `local_data`：本地覆盖数据（JSONB）
- `is_modified`：是否存在本地修改
- `last_synced_at`：最近回源时间

---

## 四、当前目录（实际）

```text
ms_tmdb/
├── backend/
│   ├── api/tmdb.api
│   ├── config/config.go
│   ├── etc/tmdb.yaml
│   ├── internal/
│   │   ├── handler/admin/admin_handler.go
│   │   ├── logic/proxy/proxy.go
│   │   ├── middleware/tmdb_proxy_middleware.go
│   │   ├── model/model.go
│   │   ├── svc/service_context.go
│   │   └── types/types.go
│   ├── pkg/
│   │   ├── tmdbclient/
│   │   └── result/
│   ├── xerr/errorx.go
│   ├── go.mod
│   └── tmdb.go
├── frontend/
├── docs/
│   ├── project-plan.md
│   ├── implementation-schedule.md
│   ├── api-reference.md
│   └── frontend-plan.md
└── docker/
    └── docker-compose.yml
```

---

## 五、已完成能力

### 5.1 基础设施

- go-zero 服务启动入口已完成（`backend/tmdb.go`）。
- PostgreSQL 连接与自动迁移已接入（`backend/internal/svc/ServiceContext` + `AutoMigrate`）。
- `docker/docker-compose.yml` 已提供 PostgreSQL/Redis。

### 5.2 TMDB 客户端

`backend/pkg/tmdbclient` 已支持：

- 电影、电视剧、人物详情与子接口
- 搜索、发现、趋势、类型列表、Find
- 电影/电视剧/人物列表接口
- 请求限流、语言/分页/区域参数透传

### 5.3 代理与缓存

- 中间件支持路径分发与参数透传。
- 三个详情接口具备 Read-Through 缓存与回源更新。

### 5.4 管理接口

已实现：

- `PUT /api/admin/movie/:id`
- `PUT /api/admin/tv/:id`
- `PUT /api/admin/person/:id`
- `POST /api/admin/sync/movie/:id`
- `POST /api/admin/sync/tv/:id`
- `POST /api/admin/sync/person/:id`
- `DELETE /api/admin/movie/:id/local`
- `DELETE /api/admin/tv/:id/local`
- `GET /api/admin/stats`

---

## 六、当前缺口

1. 测试覆盖不足：尚未建立系统化单测/集成测试。
2. 缓存模型仍较轻量：季/集/图片/视频等实体尚未落库建模。
3. 部署材料不完整：缺少 `Dockerfile` 与一键化启动说明。
4. 管理接口的输入校验与错误语义可进一步细化。

---

## 七、下一步（在当前方案上迭代）

1. 增加 `proxy` 与 `admin` 模块单元测试。
2. 优化管理接口参数校验、错误码与日志。
3. 增补部署文档与 `Dockerfile`。
4. 按业务优先级逐步扩展更多本地缓存实体，而不改变当前总体架构。
