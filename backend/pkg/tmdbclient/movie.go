package tmdbclient

import (
	"encoding/json"
	"fmt"
)

// --- 电影 API ---

// GetMovie 获取电影详情
func (c *Client) GetMovie(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d", id), opts)
}

// GetMovieCredits 获取电影演职员
func (c *Client) GetMovieCredits(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d/credits", id), opts)
}

// GetMovieImages 获取电影图片
func (c *Client) GetMovieImages(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d/images", id), opts)
}

// GetMovieVideos 获取电影视频
func (c *Client) GetMovieVideos(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d/videos", id), opts)
}

// GetMovieKeywords 获取电影关键词
func (c *Client) GetMovieKeywords(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d/keywords", id), opts)
}

// GetMovieSimilar 获取相似电影
func (c *Client) GetMovieSimilar(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d/similar", id), opts)
}

// GetMovieRecommendations 获取推荐电影
func (c *Client) GetMovieRecommendations(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d/recommendations", id), opts)
}

// GetMovieExternalIDs 获取电影外部 ID
func (c *Client) GetMovieExternalIDs(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d/external_ids", id), opts)
}

// GetMovieTranslations 获取电影翻译
func (c *Client) GetMovieTranslations(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d/translations", id), opts)
}

// GetMovieReleaseDates 获取电影上映日期
func (c *Client) GetMovieReleaseDates(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d/release_dates", id), opts)
}

// GetMovieWatchProviders 获取电影观看渠道
func (c *Client) GetMovieWatchProviders(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d/watch/providers", id), opts)
}

// GetMovieAlternativeTitles 获取电影替代标题
func (c *Client) GetMovieAlternativeTitles(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/movie/%d/alternative_titles", id), opts)
}

// --- 电影列表 API ---

// GetNowPlayingMovies 正在上映
func (c *Client) GetNowPlayingMovies(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/movie/now_playing", opts)
}

// GetPopularMovies 热门电影
func (c *Client) GetPopularMovies(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/movie/popular", opts)
}

// GetTopRatedMovies 高分电影
func (c *Client) GetTopRatedMovies(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/movie/top_rated", opts)
}

// GetUpcomingMovies 即将上映
func (c *Client) GetUpcomingMovies(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/movie/upcoming", opts)
}
