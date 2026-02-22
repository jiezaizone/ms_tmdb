import http from "./http";

export function getStats() {
  return http.get("/api/admin/stats");
}

export function syncMovie(id: number) {
  return http.post(`/api/admin/sync/movie/${id}`);
}

export function syncTV(id: number) {
  return http.post(`/api/admin/sync/tv/${id}`);
}

export function syncPerson(id: number) {
  return http.post(`/api/admin/sync/person/${id}`);
}

export function updateMovie(id: number, payload: Record<string, unknown>) {
  return http.put(`/api/admin/movie/${id}`, payload);
}

export function updateTV(id: number, payload: Record<string, unknown>) {
  return http.put(`/api/admin/tv/${id}`, payload).catch((err: any) => {
    if (String(err?.message ?? "").includes("404")) {
      return http.put(`/api/admin/tv-series/${id}`, payload);
    }
    throw err;
  });
}

export function listMovies(page = 1, pageSize = 20, keyword = "", searchMode = "contains") {
  return http.get("/api/admin/movies", { params: { page, page_size: pageSize, keyword, search_mode: searchMode } });
}

export function listTV(page = 1, pageSize = 20, keyword = "", searchMode = "contains") {
  return http
    .get("/api/admin/tv-series", { params: { page, page_size: pageSize, keyword, search_mode: searchMode } })
    .catch((err: any) => {
      if (String(err?.message ?? "").includes("404")) {
        return http.get("/api/admin/tv", { params: { page, page_size: pageSize, keyword, search_mode: searchMode } });
      }
      throw err;
    });
}

export function listPeople(page = 1, pageSize = 20) {
  return http.get("/api/admin/people", { params: { page, page_size: pageSize } });
}
