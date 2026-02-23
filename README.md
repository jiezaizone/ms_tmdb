# MS-TMDB

基于 `go-zero + PostgreSQL + Vue 3` 的 TMDB 代理与本地增强平台。

项目目标：
- 统一代理 TMDB v3 API（`/api/v3/*`）。
- 对电影/剧集/人物详情做本地 Read-Through 缓存。
- 提供管理接口与前端页面，支持本地编辑覆盖与库检索。

## 功能概览

- TMDB API 代理：电影、剧集、人物、搜索、发现、趋势、类型等。
- 本地缓存与回源：详情请求优先走本地，过期后自动回源更新。
- 管理接口：手动同步、编辑本地数据、清理本地覆盖、统计。
- 前端能力：
  - 详情页查看/编辑模式切换（电影、剧集）。
  - 本地字段保存（含类型多选）。
  - Library 卡片/表格双视图与双模式模糊搜索。

## 项目结构

```text
ms_tmdb/
├─ backend/      # go-zero 后端
├─ frontend/     # Vue 3 前端
├─ docker/       # 运行镜像与部署编排
└─ docs/         # 方案、接口、迭代文档
```

## 快速开始

### 1. 配置后端数据库与 TMDB

```bash
# 按需修改数据库地址、账号和 Tmdb.ApiKey
backend/etc/tmdb.yaml
```

### 2. 启动后端

```bash
cd backend
go run tmdb.go -f etc/tmdb.yaml
```

默认监听：`http://localhost:8888`

### 3. 启动前端

```bash
cd frontend
pnpm install
pnpm dev
```

默认访问：`http://localhost:5173`

### 4. （可选）使用 Docker 运行一体化镜像

```bash
cd docker
# 先修改 tmdb.yaml 中的数据库连接和 Tmdb.ApiKey
MS_TMDB_IMAGE=ghcr.io/<your-org>/<your-repo>:latest docker compose up -d
```

默认访问：`http://localhost:8080`

## 常用入口

- 前端页面：
  - `/` 首页
  - `/library` 本地库
  - `/movie/:id` 电影详情
  - `/tv/:id` 剧集详情
- 后端接口：
  - `/api/v3/*` TMDB 代理接口
  - `/api/admin/*` 本地管理接口

## 配置说明

- 后端配置文件：`backend/etc/tmdb.yaml`
- Docker 运行配置：`docker/tmdb.yaml`
- 关键配置项：
  - `Postgres.Host / Postgres.Port / Postgres.User / Postgres.Password / Postgres.DBName / Postgres.SSLMode`
  - `Tmdb.ApiKey`
  - `Tmdb.BaseURL`
  - `Tmdb.DefaultLanguage`

建议将 `Tmdb.ApiKey` 替换为你自己的密钥后再部署。

## 文档

- `docs/project-plan.md`
- `docs/api-reference.md`
- `docs/frontend-plan.md`
- `docs/implementation-schedule.md`
