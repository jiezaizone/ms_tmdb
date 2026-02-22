import http from "./http";

export function getPopularPeople(page = 1, language = "zh-CN") {
  return http.get("/api/v3/person/popular", { params: { page, language } });
}

export function getPersonDetail(id: number, language = "zh-CN", append = "combined_credits,images") {
  return http.get(`/api/v3/person/${id}`, { params: { language, append_to_response: append } });
}
