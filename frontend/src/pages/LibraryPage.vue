<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { listMovies, listTV } from "@/api/admin";
import { tmdbImg } from "@/api/tmdb";

type MediaTab = "movie" | "tv";
type ViewMode = "grid" | "table";
type SearchMode = "contains" | "prefix";

const activeTab = ref<MediaTab>("movie");
const viewMode = ref<ViewMode>("grid");
const searchMode = ref<SearchMode>("contains");
const keywordInput = ref("");
const keyword = ref("");
const loading = ref(false);
const error = ref("");
const items = ref<any[]>([]);
const total = ref(0);
const page = ref(1);
const pageSize = 20;

async function loadData() {
  loading.value = true;
  error.value = "";
  try {
    let resp: any;
    if (activeTab.value === "movie") resp = await listMovies(page.value, pageSize, keyword.value, searchMode.value);
    else resp = await listTV(page.value, pageSize, keyword.value, searchMode.value);
    items.value = resp.data?.results ?? [];
    total.value = resp.data?.total ?? 0;
  } catch (err: any) {
    error.value = err.message ?? "加载失败";
  } finally {
    loading.value = false;
  }
}

function switchTab(tab: MediaTab) {
  activeTab.value = tab;
  page.value = 1;
}

function applySearch() {
  keyword.value = keywordInput.value.trim();
  page.value = 1;
}

function resetSearch() {
  keywordInput.value = "";
  keyword.value = "";
  page.value = 1;
}

const totalPages = () => Math.ceil(total.value / pageSize) || 1;

function gotoPage(p: number) {
  if (p < 1 || p > totalPages()) return;
  page.value = p;
}

function routeByItem(item: any) {
  if (activeTab.value === "movie") return `/movie/${item.tmdb_id}`;
  return `/tv/${item.tmdb_id}`;
}

watch([activeTab, page, keyword, searchMode], loadData);
onMounted(loadData);
</script>

<template>
  <section class="flex flex-wrap items-center gap-2">
    <div class="flex items-center gap-2 rounded-full bg-white/70 p-1 shadow-soft">
      <button
        v-for="tab in ([
          { key: 'movie', label: '🎬 电影' },
          { key: 'tv', label: '📺 剧集' },
        ] as const)"
        :key="tab.key"
        class="rounded-full px-5 py-2 text-sm transition"
        :class="activeTab === tab.key ? 'bg-pine text-white' : 'text-ink hover:bg-sand/70'"
        @click="switchTab(tab.key as MediaTab)"
      >
        {{ tab.label }}
      </button>
    </div>

    <div class="flex items-center gap-1 rounded-full bg-white/70 p-1 shadow-soft">
      <button
        class="rounded-full px-4 py-1.5 text-xs transition"
        :class="viewMode === 'grid' ? 'bg-coral text-white' : 'text-ink hover:bg-sand/70'"
        @click="viewMode = 'grid'"
      >
        卡片
      </button>
      <button
        class="rounded-full px-4 py-1.5 text-xs transition"
        :class="viewMode === 'table' ? 'bg-coral text-white' : 'text-ink hover:bg-sand/70'"
        @click="viewMode = 'table'"
      >
        表格
      </button>
    </div>
  </section>

  <section class="card mt-4">
    <div class="grid gap-3 md:grid-cols-[1fr_auto_auto_auto] md:items-center">
      <input
        v-model="keywordInput"
        class="w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
        placeholder="输入片名/剧名关键词"
        @keyup.enter="applySearch"
      />
      <select
        v-model="searchMode"
        class="rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
      >
        <option value="contains">模糊包含</option>
        <option value="prefix">前缀匹配</option>
      </select>
      <button class="rounded-lg bg-pine px-4 py-2 text-sm text-white hover:bg-pine/90" @click="applySearch">
        搜索
      </button>
      <button class="rounded-lg border border-black/10 bg-white px-4 py-2 text-sm hover:bg-sand/50" @click="resetSearch">
        重置
      </button>
    </div>
  </section>

  <p v-if="loading" class="card mt-4 text-sm text-black/60">加载中...</p>
  <p v-else-if="error" class="card mt-4 text-sm text-red-600">{{ error }}</p>

  <template v-else>
    <section class="mt-4 flex items-center justify-between">
      <p class="text-sm text-black/60">
        共 <strong>{{ total }}</strong> 条记录 · 第 {{ page }}/{{ totalPages() }} 页
      </p>
    </section>

    <section v-if="viewMode === 'grid'" class="mt-4 poster-grid">
      <RouterLink
        v-for="item in items"
        :key="item.tmdb_id"
        :to="routeByItem(item)"
        class="poster-card"
      >
        <img
          :src="tmdbImg(item.poster_path, 'w185')"
          :alt="item.title || item.name"
          class="poster-img"
          loading="lazy"
        />
        <div class="poster-info">
          <p class="truncate text-sm font-medium">{{ item.title || item.name }}</p>
          <p class="text-xs text-black/55">
            ⭐ {{ (item.vote_average ?? 0).toFixed(1) }}
            <span class="ml-1">{{ (item.release_date || item.first_air_date || "").slice(0, 4) }}</span>
          </p>
          <span v-if="item.is_modified" class="mt-1 inline-block rounded bg-coral/20 px-1.5 py-0.5 text-[10px] text-coral">
            已修改
          </span>
        </div>
      </RouterLink>
    </section>

    <section v-else class="card mt-4 overflow-x-auto p-0">
      <table class="min-w-full text-left text-sm">
        <thead class="bg-black/[0.04] text-xs uppercase tracking-wide text-black/60">
          <tr>
            <th class="px-4 py-3">TMDB ID</th>
            <th class="px-4 py-3">名称</th>
            <th class="px-4 py-3">评分</th>
            <th class="px-4 py-3">日期</th>
            <th class="px-4 py-3">热度</th>
            <th class="px-4 py-3">状态</th>
            <th class="px-4 py-3">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="item in items"
            :key="item.tmdb_id"
            class="border-t border-black/5 hover:bg-black/[0.02]"
          >
            <td class="px-4 py-3">{{ item.tmdb_id }}</td>
            <td class="px-4 py-3">
              <p class="font-medium">{{ item.title || item.name }}</p>
              <p class="text-xs text-black/50">{{ item.original_title || item.original_name }}</p>
            </td>
            <td class="px-4 py-3">⭐ {{ (item.vote_average ?? 0).toFixed(1) }}</td>
            <td class="px-4 py-3">{{ item.release_date || item.first_air_date || "-" }}</td>
            <td class="px-4 py-3">{{ (item.popularity ?? 0).toFixed(1) }}</td>
            <td class="px-4 py-3">
              <span
                v-if="item.is_modified"
                class="inline-block rounded bg-coral/20 px-2 py-0.5 text-[11px] text-coral"
              >
                已修改
              </span>
              <span v-else class="text-xs text-black/45">未修改</span>
            </td>
            <td class="px-4 py-3">
              <RouterLink :to="routeByItem(item)" class="text-pine hover:underline">
                查看详情
              </RouterLink>
            </td>
          </tr>
          <tr v-if="items.length === 0">
            <td colspan="7" class="px-4 py-8 text-center text-black/50">无数据</td>
          </tr>
        </tbody>
      </table>
    </section>

    <section class="mt-6 flex items-center justify-center gap-2">
      <button
        class="rounded-lg border border-black/10 bg-white px-3 py-1.5 text-sm hover:bg-sand/50 disabled:opacity-40"
        :disabled="page <= 1"
        @click="gotoPage(page - 1)"
      >
        上一页
      </button>
      <span class="px-3 text-sm text-black/60">{{ page }} / {{ totalPages() }}</span>
      <button
        class="rounded-lg border border-black/10 bg-white px-3 py-1.5 text-sm hover:bg-sand/50 disabled:opacity-40"
        :disabled="page >= totalPages()"
        @click="gotoPage(page + 1)"
      >
        下一页
      </button>
    </section>
  </template>
</template>
