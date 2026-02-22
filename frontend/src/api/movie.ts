import http from "./http";

export function getPopularMovies(page = 1, language = "zh-CN") {
  return http.get("/api/v3/movie/popular", { params: { page, language } });
}

export function getMovieDetail(id: number, language = "zh-CN", append = "credits,videos,images") {
  return http.get(`/api/v3/movie/${id}`, { params: { language, append_to_response: append } });
}

export function getMovieGenreList(language = "zh-CN") {
  return http.get("/api/v3/genre/movie/list", { params: { language } });
}
