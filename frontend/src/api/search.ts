import http from "./http";

export type SearchType = "movie" | "tv" | "person" | "multi";

export function searchByType(type: SearchType, query: string, page = 1, language = "zh-CN") {
  return http.get(`/api/v3/search/${type}`, {
    params: { query, page, language },
  });
}
