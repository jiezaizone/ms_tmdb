<script setup lang="ts">
import { computed, onMounted, ref, watch } from "vue";
import { useRoute } from "vue-router";
import { updateTV } from "@/api/admin";
import { getTVDetail, getTVGenreList } from "@/api/tv";
import { tmdbImg } from "@/api/tmdb";

type GenreOption = {
  id: number;
  name: string;
};

type TVEditForm = {
  name: string;
  original_name: string;
  genre_names: string[];
  type: string;
  tagline: string;
  first_air_date: string;
  status: string;
  number_of_seasons: string;
  number_of_episodes: string;
  original_language: string;
  homepage: string;
  poster_path: string;
  backdrop_path: string;
  vote_average: string;
  popularity: string;
  overview: string;
};

const route = useRoute();
const loading = ref(false);
const error = ref("");
const detail = ref<any>(null);
const isEditing = ref(false);
const saving = ref(false);
const saveError = ref("");
const saveMessage = ref("");
const genreOptions = ref<GenreOption[]>([]);
const genreKeyword = ref("");
const filteredGenreOptions = computed(() => {
  const keyword = genreKeyword.value.trim().toLowerCase();
  if (!keyword) {
    return genreOptions.value;
  }
  return genreOptions.value.filter((genre) => genre.name.toLowerCase().includes(keyword));
});
const editForm = ref<TVEditForm>({
  name: "",
  original_name: "",
  genre_names: [],
  type: "",
  tagline: "",
  first_air_date: "",
  status: "",
  number_of_seasons: "",
  number_of_episodes: "",
  original_language: "",
  homepage: "",
  poster_path: "",
  backdrop_path: "",
  vote_average: "",
  popularity: "",
  overview: "",
});

const tvId = computed(() => Number(route.params.id));

function resetEditForm(data: any) {
  editForm.value = {
    name: data?.name ?? "",
    original_name: data?.original_name ?? "",
    genre_names: Array.isArray(data?.genres) ? data.genres.map((g: any) => String(g?.name ?? "").trim()).filter(Boolean) : [],
    type: data?.type ?? "",
    tagline: data?.tagline ?? "",
    first_air_date: data?.first_air_date ?? "",
    status: data?.status ?? "",
    number_of_seasons: data?.number_of_seasons != null ? String(data.number_of_seasons) : "",
    number_of_episodes: data?.number_of_episodes != null ? String(data.number_of_episodes) : "",
    original_language: data?.original_language ?? "",
    homepage: data?.homepage ?? "",
    poster_path: data?.poster_path ?? "",
    backdrop_path: data?.backdrop_path ?? "",
    vote_average: data?.vote_average != null ? String(data.vote_average) : "",
    popularity: data?.popularity != null ? String(data.popularity) : "",
    overview: data?.overview ?? "",
  };
}

function normalizeGenreOptions(raw: any): GenreOption[] {
  if (!Array.isArray(raw)) return [];
  return raw
    .map((item: any, idx: number) => ({
      id: Number(item?.id) || idx + 1,
      name: String(item?.name ?? "").trim(),
    }))
    .filter((item: GenreOption) => !!item.name);
}

async function loadGenreOptions() {
  try {
    const resp = await getTVGenreList();
    const options = normalizeGenreOptions(resp.data?.genres);
    if (options.length > 0) {
      genreOptions.value = options;
      return;
    }
  } catch {
    // 忽略类型列表加载失败，降级使用详情已有类型
  }

  genreOptions.value = normalizeGenreOptions(detail.value?.genres);
}

function enterEditMode() {
  if (!detail.value) return;
  resetEditForm(detail.value);
  genreKeyword.value = "";
  saveError.value = "";
  saveMessage.value = "";
  isEditing.value = true;
}

function cancelEditMode() {
  if (detail.value) {
    resetEditForm(detail.value);
  }
  genreKeyword.value = "";
  saveError.value = "";
  isEditing.value = false;
}

async function loadData() {
  if (!tvId.value) {
    error.value = "无效剧集 ID";
    return;
  }
  loading.value = true;
  error.value = "";
  try {
    const resp = await getTVDetail(tvId.value);
    detail.value = resp.data;
    resetEditForm(resp.data);
    await loadGenreOptions();
    genreKeyword.value = "";
    isEditing.value = false;
  } catch (err: any) {
    error.value = err.message ?? "加载失败";
  } finally {
    loading.value = false;
  }
}

function parseOptionalInt(raw: string): number | undefined {
  const text = raw.trim();
  if (!text) return undefined;
  const value = Number(text);
  if (!Number.isFinite(value)) return undefined;
  return Math.trunc(value);
}

function parseOptionalFloat(raw: string): number | undefined {
  const text = raw.trim();
  if (!text) return undefined;
  const value = Number(text);
  if (!Number.isFinite(value)) return undefined;
  return value;
}

async function saveTVChanges() {
  if (!tvId.value) {
    saveError.value = "无效剧集 ID";
    return;
  }
  const seasons = parseOptionalInt(editForm.value.number_of_seasons);
  if (editForm.value.number_of_seasons.trim() && seasons === undefined) {
    saveError.value = "季数必须是数字";
    return;
  }
  const episodes = parseOptionalInt(editForm.value.number_of_episodes);
  if (editForm.value.number_of_episodes.trim() && episodes === undefined) {
    saveError.value = "集数必须是数字";
    return;
  }
  const voteAverage = parseOptionalFloat(editForm.value.vote_average);
  if (editForm.value.vote_average.trim() && voteAverage === undefined) {
    saveError.value = "评分必须是数字";
    return;
  }
  const popularity = parseOptionalFloat(editForm.value.popularity);
  if (editForm.value.popularity.trim() && popularity === undefined) {
    saveError.value = "热度必须是数字";
    return;
  }

  saving.value = true;
  saveError.value = "";
  saveMessage.value = "";
  try {
    const payload: Record<string, unknown> = {
      name: editForm.value.name.trim(),
      original_name: editForm.value.original_name.trim(),
      genre_names: editForm.value.genre_names,
      type: editForm.value.type.trim(),
      tagline: editForm.value.tagline.trim(),
      first_air_date: editForm.value.first_air_date.trim(),
      status: editForm.value.status.trim(),
      original_language: editForm.value.original_language.trim(),
      homepage: editForm.value.homepage.trim(),
      poster_path: editForm.value.poster_path.trim(),
      backdrop_path: editForm.value.backdrop_path.trim(),
      overview: editForm.value.overview.trim(),
    };
    if (seasons !== undefined) {
      payload.number_of_seasons = seasons;
    }
    if (episodes !== undefined) {
      payload.number_of_episodes = episodes;
    }
    if (voteAverage !== undefined) {
      payload.vote_average = voteAverage;
    }
    if (popularity !== undefined) {
      payload.popularity = popularity;
    }

    await updateTV(tvId.value, payload);
    saveMessage.value = "已保存到本地数据库";
    isEditing.value = false;
    await loadData();
  } catch (err: any) {
    saveError.value = err.message ?? "保存失败";
  } finally {
    saving.value = false;
  }
}

onMounted(loadData);
watch(tvId, loadData);
</script>

<template>
  <p v-if="loading" class="card text-sm text-black/60">加载中...</p>
  <p v-else-if="error" class="card text-sm text-red-600">{{ error }}</p>

  <template v-else-if="detail">
    <!-- 背景横幅 -->
    <section
      class="hero-banner"
      :style="{ backgroundImage: `url(${tmdbImg(detail.backdrop_path, 'w780')})` }"
    >
      <div class="hero-overlay">
        <h1 class="text-2xl font-bold text-white md:text-3xl">{{ detail.name || detail.original_name }}</h1>
        <p class="mt-1 text-sm text-white/70">{{ detail.tagline }}</p>
      </div>
    </section>

    <section class="card mt-4">
      <div class="detail-layout">
        <div class="detail-poster">
          <img
            :src="tmdbImg(detail.poster_path, 'w342')"
            :alt="detail.name"
            class="w-full rounded-xl shadow-soft"
          />
        </div>

        <div class="detail-info">
          <h2 class="text-xl font-bold">{{ detail.name }}</h2>
          <p v-if="detail.original_name !== detail.name" class="text-sm text-black/55">
            {{ detail.original_name }}
          </p>

          <div class="mt-3 flex flex-wrap gap-2">
            <span class="badge">⭐ {{ detail.vote_average?.toFixed(1) ?? "-" }}</span>
            <span class="badge">📅 {{ detail.first_air_date ?? "-" }}</span>
            <span v-if="detail.number_of_seasons" class="badge">
              {{ detail.number_of_seasons }} 季 · {{ detail.number_of_episodes }} 集
            </span>
            <span class="badge">{{ detail.status ?? "未知" }}</span>
          </div>

          <div v-if="detail.genres?.length" class="mt-3 flex flex-wrap gap-1.5">
            <span
              v-for="g in detail.genres"
              :key="g.id"
              class="rounded-full bg-sand/60 px-3 py-1 text-xs text-ink"
            >
              {{ g.name }}
            </span>
          </div>

          <p class="mt-4 text-sm leading-relaxed text-black/75">
            {{ detail.overview || "暂无简介" }}
          </p>

          <div class="mt-6 rounded-xl border border-black/10 bg-white/70 p-4">
            <div class="flex items-center justify-between gap-3">
              <h3 class="text-sm font-semibold">本地信息编辑</h3>
              <button
                v-if="!isEditing"
                class="rounded-lg border border-black/10 bg-white px-3 py-1.5 text-xs hover:bg-sand/50"
                @click="enterEditMode"
              >
                编辑
              </button>
            </div>

            <p v-if="!isEditing" class="mt-2 text-xs text-black/60">
              当前为查看模式，点击“编辑”后可修改并保存到本地数据库。
            </p>

            <div v-else class="mt-3">
              <div class="grid gap-3 md:grid-cols-2">
                <label class="text-xs text-black/60">
                  剧名
                  <input
                    v-model="editForm.name"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="剧集标题"
                  />
                </label>
                <label class="text-xs text-black/60">
                  原始剧名
                  <input
                    v-model="editForm.original_name"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="Original Name"
                  />
                </label>
                <label class="text-xs text-black/60 md:col-span-2">
                  类型（多选）
                  <div class="mt-1 flex flex-wrap gap-2 rounded-lg border border-black/10 bg-white p-2">
                    <input
                      v-model="genreKeyword"
                      class="w-full rounded-md border border-black/10 bg-sand/20 px-2.5 py-1.5 text-xs"
                      placeholder="筛选类型"
                    />
                    <label
                      v-for="genre in filteredGenreOptions"
                      :key="genre.id"
                      class="inline-flex items-center gap-1.5 rounded-md border border-black/10 px-2 py-1 text-xs"
                    >
                      <input
                        v-model="editForm.genre_names"
                        type="checkbox"
                        :value="genre.name"
                      />
                      <span>{{ genre.name }}</span>
                    </label>
                    <span v-if="!genreOptions.length" class="px-1 py-1 text-xs text-black/50">
                      暂无可选类型
                    </span>
                    <span v-else-if="!filteredGenreOptions.length" class="px-1 py-1 text-xs text-black/50">
                      无匹配类型
                    </span>
                  </div>
                </label>
                <label class="text-xs text-black/60">
                  首播日期
                  <input
                    v-model="editForm.first_air_date"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="YYYY-MM-DD"
                  />
                </label>
                <label class="text-xs text-black/60">
                  状态
                  <input
                    v-model="editForm.status"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="Returning Series"
                  />
                </label>
                <label class="text-xs text-black/60">
                  剧集类型
                  <input
                    v-model="editForm.type"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="Scripted / Miniseries"
                  />
                </label>
                <label class="text-xs text-black/60">
                  季数
                  <input
                    v-model="editForm.number_of_seasons"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="Seasons"
                  />
                </label>
                <label class="text-xs text-black/60">
                  集数
                  <input
                    v-model="editForm.number_of_episodes"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="Episodes"
                  />
                </label>
                <label class="text-xs text-black/60">
                  原始语言
                  <input
                    v-model="editForm.original_language"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="zh / en"
                  />
                </label>
                <label class="text-xs text-black/60">
                  主页链接
                  <input
                    v-model="editForm.homepage"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="https://..."
                  />
                </label>
                <label class="text-xs text-black/60">
                  海报路径
                  <input
                    v-model="editForm.poster_path"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="/poster.jpg"
                  />
                </label>
                <label class="text-xs text-black/60">
                  背景图路径
                  <input
                    v-model="editForm.backdrop_path"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="/backdrop.jpg"
                  />
                </label>
                <label class="text-xs text-black/60">
                  评分
                  <input
                    v-model="editForm.vote_average"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="8.4"
                  />
                </label>
                <label class="text-xs text-black/60">
                  热度
                  <input
                    v-model="editForm.popularity"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="210.5"
                  />
                </label>
                <label class="text-xs text-black/60 md:col-span-2">
                  简介
                  <textarea
                    v-model="editForm.overview"
                    rows="4"
                    class="mt-1 w-full rounded-lg border border-black/10 bg-white px-3 py-2 text-sm"
                    placeholder="简介"
                  />
                </label>
              </div>

              <div class="mt-3 flex items-center gap-3">
                <button
                  class="rounded-lg bg-coral px-4 py-2 text-sm font-medium text-white hover:bg-coral/90 disabled:opacity-60"
                  :disabled="saving"
                  @click="saveTVChanges"
                >
                  {{ saving ? "保存中..." : "保存到本地数据库" }}
                </button>
                <button
                  class="rounded-lg border border-black/10 bg-white px-4 py-2 text-sm hover:bg-sand/50 disabled:opacity-60"
                  :disabled="saving"
                  @click="cancelEditMode"
                >
                  取消
                </button>
              </div>
            </div>

            <div class="mt-2">
              <span v-if="saveMessage" class="text-xs text-green-700">{{ saveMessage }}</span>
              <span v-if="saveError" class="text-xs text-red-600">{{ saveError }}</span>
            </div>
          </div>

          <!-- 季列表 -->
          <div v-if="detail.seasons?.length" class="mt-6">
            <h3 class="mb-2 text-sm font-semibold">季列表</h3>
            <div class="cast-grid">
              <div v-for="s in detail.seasons" :key="s.id" class="cast-card">
                <img
                  :src="tmdbImg(s.poster_path, 'w185')"
                  :alt="s.name"
                  class="cast-img"
                  loading="lazy"
                />
                <p class="mt-1 truncate text-xs font-medium">{{ s.name }}</p>
                <p class="truncate text-xs text-black/50">{{ s.episode_count }} 集</p>
              </div>
            </div>
          </div>

          <!-- 演员 -->
          <div v-if="detail.credits?.cast?.length" class="mt-6">
            <h3 class="mb-2 text-sm font-semibold">主要演员</h3>
            <div class="cast-grid">
              <div v-for="c in detail.credits.cast.slice(0, 8)" :key="c.id" class="cast-card">
                <RouterLink :to="`/person/${c.id}`">
                  <img
                    :src="tmdbImg(c.profile_path, 'w185')"
                    :alt="c.name"
                    class="cast-img"
                    loading="lazy"
                  />
                </RouterLink>
                <p class="mt-1 truncate text-xs font-medium">{{ c.name }}</p>
                <p class="truncate text-xs text-black/50">{{ c.character }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </template>
</template>
