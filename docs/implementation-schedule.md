# MS-TMDB 实施计划（进度同步）

更新时间：2026-02-21

## 当前状态总览

- 当前服务可编译、可运行，在 `backend/` 下执行 `go test ./...` 与 `go build ./...` 通过（默认 Go 缓存路径）。
- 主体模式为“TMDB 透明代理 + 部分本地缓存（电影/电视剧/人物详情）+ 管理接口”。
- 文档已从“纯规划态”更新为“已落地能力 + 待办事项”。

---

## 阶段进度

### 阶段一：项目初始化与基础设施

- [x] 初始化 Go 模块（`backend/go.mod`）
- [x] 搭建 go-zero 服务入口与目录结构
- [ ] `Makefile` 构建脚本
- [x] `docker/docker-compose.yml`（PostgreSQL + Redis）
- [x] 服务配置文件（`backend/etc/tmdb.yaml`）
- [x] PostgreSQL 连接初始化（`backend/internal/svc/ServiceContext`）
- [x] 自动建表迁移（`backend/internal/model/AutoMigrate`）
- [x] 统一错误码（`backend/xerr`）
- [x] 统一响应包（`backend/pkg/result`）

状态：**已完成（基础可用）**

### 阶段二：TMDB 客户端封装

- [x] 基础 HTTP 客户端、限流、参数构建（`backend/pkg/tmdbclient/client.go`）
- [x] 电影接口封装（详情、子接口、列表）
- [x] 电视剧接口封装（详情、季/集、子接口、列表）
- [x] 人物、搜索、发现、趋势、类型、Find 接口封装
- [ ] 重试机制（当前未实现）
- [ ] 强类型响应结构（当前以 `json.RawMessage` 为主）

状态：**进行中（核心能力已可用）**

### 阶段三：数据模型层开发

- [x] `Movie` / `TVSeries` / `Person` 三个核心模型
- [x] `RawJSON`（JSONB）类型封装
- [x] Read-Through 过期判断与 Upsert（详情接口）
- [x] 本地覆盖合并逻辑（管理接口中的浅合并）
- [ ] 季/集/关键词/视频/图片等扩展模型

状态：**进行中**

### 阶段四：API 定义与代码生成

- [x] `backend/api/tmdb.api` 已定义电影/电视剧/人物/搜索/发现/趋势/管理等路由
- [x] `backend/internal/types/types.go` 已生成
- [ ] 完整按 goctl 生成并接入所有 handler/logic（当前未采用该路径）

状态：**进行中（当前实现以中间件代理为主）**

### 阶段五：核心业务逻辑

- [x] 电影详情 Read-Through（本地命中/过期回源/写库）
- [x] 电视剧详情 Read-Through
- [x] 人物详情 Read-Through
- [x] 其余 TMDB v3 路径透传（中间件分发）
- [ ] 细粒度业务逻辑拆分到独立 logic 层（当前集中在 middleware/proxy）

状态：**进行中**

### 阶段六：管理接口

- [x] 修改本地数据（movie/tv/person）
- [x] 强制同步（movie/tv/person）
- [x] 清除本地修改（movie/tv）
- [x] 统计接口（`/api/admin/stats`）

状态：**已完成（当前范围）**

### 阶段七：测试与文档

- [x] 文档已同步到当前代码实现
- [ ] 单元测试（当前基本无测试文件）
- [ ] 集成测试与端到端测试

状态：**进行中**

### 阶段八：部署与优化

- [x] `docker-compose` 已提供
- [ ] `Dockerfile`
- [ ] CI/CD、监控、性能压测

状态：**未开始/部分完成**

---

## 近期优先级（建议）

1. 增加单元测试：优先 `backend/internal/logic/proxy` 与 `backend/internal/handler/admin`。
2. 抽离管理接口公共逻辑，补充严格参数校验与事务错误处理。
3. 保持当前中间件代理实现，补齐文档与代码映射关系，避免“api 定义与运行时实现”双轨漂移。
4. 增补 `Dockerfile` 与启动说明，降低环境搭建成本。
