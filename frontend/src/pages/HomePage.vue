<script setup lang="ts">
import { onMounted, ref } from "vue";
import { getPopularMovies } from "@/api/movie";
import { getPopularTV } from "@/api/tv";
import { tmdbImg } from "@/api/tmdb";

const loading = ref(false);
const error = ref("");
const movies = ref<any[]>([]);
const tvSeries = ref<any[]>([]);

async function loadData() {
  loading.value = true;
  error.value = "";
  try {
    const [movieResp, tvResp] = await Promise.all([
      getPopularMovies(1),
      getPopularTV(1),
    ]);
    movies.value = movieResp.data?.results ?? [];
    tvSeries.value = tvResp.data?.results ?? [];
  } catch (err: any) {
    error.value = err.message ?? "加载失败";
  } finally {
    loading.value = false;
  }
}

onMounted(loadData);
</script>

<template>
  <!-- 顶部横幅：用第一部热门电影的背景图 -->
  <section
    v-if="movies.length"
    class="hero-banner"
    :style="{ backgroundImage: `url(${tmdbImg(movies[0]?.backdrop_path, 'w780')})` }"
  >
    <div class="hero-overlay">
      <h2 class="text-2xl font-bold text-white md:text-3xl">{{ movies[0]?.title }}</h2>
      <p class="mt-2 max-w-xl text-sm text-white/80 line-clamp-2">{{ movies[0]?.overview }}</p>
      <RouterLink
        :to="`/movie/${movies[0]?.id}`"
        class="mt-3 inline-block rounded-full bg-coral px-5 py-2 text-sm font-medium text-white hover:bg-coral/90"
      >
        查看详情
      </RouterLink>
    </div>
  </section>

  <section class="mt-4 flex items-center justify-between">
    <p class="text-xs uppercase tracking-[0.25em] text-pine/80">今日看点</p>
    <button
      class="rounded-xl bg-pine px-4 py-2 text-sm font-medium text-white hover:bg-pine/90"
      :disabled="loading"
      @click="loadData"
    >
      {{ loading ? "刷新中..." : "刷新数据" }}
    </button>
  </section>
  <p v-if="error" class="mt-3 text-sm text-red-600">{{ error }}</p>

  <!-- 热门电影 -->
  <section class="mt-6">
    <h3 class="mb-3 text-base font-semibold">🎬 热门电影</h3>
    <div class="poster-grid">
      <RouterLink
        v-for="item in movies.slice(0, 10)"
        :key="item.id"
        :to="`/movie/${item.id}`"
        class="poster-card"
      >
        <img
          :src="tmdbImg(item.poster_path, 'w185')"
          :alt="item.title"
          class="poster-img"
          loading="lazy"
        />
        <div class="poster-info">
          <p class="truncate text-sm font-medium">{{ item.title || item.original_title }}</p>
          <p class="text-xs text-black/55">
            ⭐ {{ item.vote_average?.toFixed(1) ?? "-" }}
            <span class="ml-1">{{ item.release_date?.slice(0, 4) ?? "" }}</span>
          </p>
        </div>
      </RouterLink>
    </div>
  </section>

  <!-- 热门剧集 -->
  <section class="mt-8">
    <h3 class="mb-3 text-base font-semibold">📺 热门剧集</h3>
    <div class="poster-grid">
      <RouterLink
        v-for="item in tvSeries.slice(0, 10)"
        :key="item.id"
        :to="`/tv/${item.id}`"
        class="poster-card"
      >
        <img
          :src="tmdbImg(item.poster_path, 'w185')"
          :alt="item.name"
          class="poster-img"
          loading="lazy"
        />
        <div class="poster-info">
          <p class="truncate text-sm font-medium">{{ item.name || item.original_name }}</p>
          <p class="text-xs text-black/55">
            ⭐ {{ item.vote_average?.toFixed(1) ?? "-" }}
            <span class="ml-1">{{ item.first_air_date?.slice(0, 4) ?? "" }}</span>
          </p>
        </div>
      </RouterLink>
    </div>
  </section>
</template>
