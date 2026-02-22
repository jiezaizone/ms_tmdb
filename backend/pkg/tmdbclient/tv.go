package tmdbclient

import (
	"encoding/json"
	"fmt"
)

// --- 电视剧 API ---

// GetTVSeries 获取电视剧详情
func (c *Client) GetTVSeries(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d", id), opts)
}

// GetTVSeriesCredits 获取电视剧演职员
func (c *Client) GetTVSeriesCredits(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/credits", id), opts)
}

// GetTVSeriesAggregateCredits 获取电视剧汇总演职员
func (c *Client) GetTVSeriesAggregateCredits(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/aggregate_credits", id), opts)
}

// GetTVSeriesImages 获取电视剧图片
func (c *Client) GetTVSeriesImages(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/images", id), opts)
}

// GetTVSeriesVideos 获取电视剧视频
func (c *Client) GetTVSeriesVideos(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/videos", id), opts)
}

// GetTVSeriesKeywords 获取电视剧关键词
func (c *Client) GetTVSeriesKeywords(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/keywords", id), opts)
}

// GetTVSeriesSimilar 获取相似电视剧
func (c *Client) GetTVSeriesSimilar(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/similar", id), opts)
}

// GetTVSeriesRecommendations 获取推荐电视剧
func (c *Client) GetTVSeriesRecommendations(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/recommendations", id), opts)
}

// GetTVSeriesExternalIDs 获取电视剧外部 ID
func (c *Client) GetTVSeriesExternalIDs(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/external_ids", id), opts)
}

// GetTVSeriesContentRatings 获取电视剧内容分级
func (c *Client) GetTVSeriesContentRatings(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/content_ratings", id), opts)
}

// GetTVSeriesTranslations 获取电视剧翻译
func (c *Client) GetTVSeriesTranslations(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/translations", id), opts)
}

// GetTVSeriesWatchProviders 获取电视剧观看渠道
func (c *Client) GetTVSeriesWatchProviders(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/watch/providers", id), opts)
}

// --- 电视剧季 API ---

// GetTVSeason 获取电视剧季详情
func (c *Client) GetTVSeason(seriesID, seasonNumber int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/season/%d", seriesID, seasonNumber), opts)
}

// GetTVSeasonCredits 获取季演职员
func (c *Client) GetTVSeasonCredits(seriesID, seasonNumber int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/season/%d/credits", seriesID, seasonNumber), opts)
}

// GetTVSeasonImages 获取季图片
func (c *Client) GetTVSeasonImages(seriesID, seasonNumber int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/season/%d/images", seriesID, seasonNumber), opts)
}

// GetTVSeasonVideos 获取季视频
func (c *Client) GetTVSeasonVideos(seriesID, seasonNumber int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/season/%d/videos", seriesID, seasonNumber), opts)
}

// --- 电视剧集 API ---

// GetTVEpisode 获取电视剧集详情
func (c *Client) GetTVEpisode(seriesID, seasonNumber, episodeNumber int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/season/%d/episode/%d", seriesID, seasonNumber, episodeNumber), opts)
}

// GetTVEpisodeCredits 获取集演职员
func (c *Client) GetTVEpisodeCredits(seriesID, seasonNumber, episodeNumber int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/season/%d/episode/%d/credits", seriesID, seasonNumber, episodeNumber), opts)
}

// GetTVEpisodeImages 获取集图片
func (c *Client) GetTVEpisodeImages(seriesID, seasonNumber, episodeNumber int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/tv/%d/season/%d/episode/%d/images", seriesID, seasonNumber, episodeNumber), opts)
}

// --- 电视剧列表 API ---

// GetAiringTodayTV 今日播出
func (c *Client) GetAiringTodayTV(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/tv/airing_today", opts)
}

// GetOnTheAirTV 正在播出
func (c *Client) GetOnTheAirTV(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/tv/on_the_air", opts)
}

// GetPopularTV 热门电视剧
func (c *Client) GetPopularTV(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/tv/popular", opts)
}

// GetTopRatedTV 高分电视剧
func (c *Client) GetTopRatedTV(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/tv/top_rated", opts)
}
