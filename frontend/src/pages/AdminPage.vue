<script setup lang="ts">
import { ref } from "vue";
import { getStats, syncMovie, syncPerson, syncTV } from "@/api/admin";

const stats = ref<Record<string, number> | null>(null);
const syncType = ref<"movie" | "tv" | "person">("movie");
const targetId = ref<number | null>(null);
const loading = ref(false);
const message = ref("");
const error = ref("");

async function loadStats() {
  loading.value = true;
  error.value = "";
  try {
    const resp = await getStats();
    stats.value = resp.data;
  } catch (err: any) {
    error.value = err.message ?? "读取统计失败";
  } finally {
    loading.value = false;
  }
}

async function doSync() {
  if (!targetId.value) {
    error.value = "请输入目标 ID";
    return;
  }
  loading.value = true;
  error.value = "";
  message.value = "";
  try {
    if (syncType.value === "movie") await syncMovie(targetId.value);
    if (syncType.value === "tv") await syncTV(targetId.value);
    if (syncType.value === "person") await syncPerson(targetId.value);
    message.value = `同步成功：${syncType.value}/${targetId.value}`;
    await loadStats();
  } catch (err: any) {
    error.value = err.message ?? "同步失败";
  } finally {
    loading.value = false;
  }
}

loadStats();
</script>

<template>
  <section class="grid gap-4 md:grid-cols-2">
    <article class="card">
      <h2 class="mb-3 text-lg font-semibold">数据统计</h2>
      <button class="rounded-xl bg-pine px-4 py-2 text-sm font-medium text-white hover:bg-pine/90" @click="loadStats">
        刷新统计
      </button>
      <pre class="mt-3 overflow-auto rounded-xl bg-black/90 p-3 text-xs text-green-300">{{ stats }}</pre>
    </article>

    <article class="card">
      <h2 class="mb-3 text-lg font-semibold">手动同步</h2>
      <div class="grid gap-3">
        <select v-model="syncType" class="rounded-xl border border-black/10 bg-white px-3 py-2">
          <option value="movie">movie</option>
          <option value="tv">tv</option>
          <option value="person">person</option>
        </select>
        <input
          v-model.number="targetId"
          type="number"
          min="1"
          class="rounded-xl border border-black/10 bg-white px-3 py-2"
          placeholder="输入 TMDB ID"
        />
        <button class="rounded-xl bg-coral px-4 py-2 font-medium text-white hover:bg-coral/90" @click="doSync">
          {{ loading ? "处理中..." : "执行同步" }}
        </button>
      </div>
      <p v-if="message" class="mt-3 text-sm text-green-700">{{ message }}</p>
      <p v-if="error" class="mt-3 text-sm text-red-600">{{ error }}</p>
    </article>
  </section>
</template>
