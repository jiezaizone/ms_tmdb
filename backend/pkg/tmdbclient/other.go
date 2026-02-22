package tmdbclient

import (
	"encoding/json"
	"fmt"
)

// --- 人物 API ---

// GetPerson 获取人物详情
func (c *Client) GetPerson(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/person/%d", id), opts)
}

// GetPersonMovieCredits 获取人物电影作品
func (c *Client) GetPersonMovieCredits(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/person/%d/movie_credits", id), opts)
}

// GetPersonTVCredits 获取人物电视剧作品
func (c *Client) GetPersonTVCredits(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/person/%d/tv_credits", id), opts)
}

// GetPersonCombinedCredits 获取人物全部作品
func (c *Client) GetPersonCombinedCredits(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/person/%d/combined_credits", id), opts)
}

// GetPersonImages 获取人物图片
func (c *Client) GetPersonImages(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/person/%d/images", id), opts)
}

// GetPersonExternalIDs 获取人物外部 ID
func (c *Client) GetPersonExternalIDs(id int, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/person/%d/external_ids", id), opts)
}

// GetPopularPeople 获取热门人物
func (c *Client) GetPopularPeople(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/person/popular", opts)
}

// --- 搜索 API ---

// SearchMovie 搜索电影
func (c *Client) SearchMovie(query string, opts *RequestOption) (json.RawMessage, error) {
	if opts == nil {
		opts = &RequestOption{}
	}
	if opts.ExtraParams == nil {
		opts.ExtraParams = make(map[string]string)
	}
	opts.ExtraParams["query"] = query
	return c.Get("/search/movie", opts)
}

// SearchTV 搜索电视剧
func (c *Client) SearchTV(query string, opts *RequestOption) (json.RawMessage, error) {
	if opts == nil {
		opts = &RequestOption{}
	}
	if opts.ExtraParams == nil {
		opts.ExtraParams = make(map[string]string)
	}
	opts.ExtraParams["query"] = query
	return c.Get("/search/tv", opts)
}

// SearchPerson 搜索人物
func (c *Client) SearchPerson(query string, opts *RequestOption) (json.RawMessage, error) {
	if opts == nil {
		opts = &RequestOption{}
	}
	if opts.ExtraParams == nil {
		opts.ExtraParams = make(map[string]string)
	}
	opts.ExtraParams["query"] = query
	return c.Get("/search/person", opts)
}

// SearchMulti 多类型搜索
func (c *Client) SearchMulti(query string, opts *RequestOption) (json.RawMessage, error) {
	if opts == nil {
		opts = &RequestOption{}
	}
	if opts.ExtraParams == nil {
		opts.ExtraParams = make(map[string]string)
	}
	opts.ExtraParams["query"] = query
	return c.Get("/search/multi", opts)
}

// SearchKeyword 搜索关键词
func (c *Client) SearchKeyword(query string, opts *RequestOption) (json.RawMessage, error) {
	if opts == nil {
		opts = &RequestOption{}
	}
	if opts.ExtraParams == nil {
		opts.ExtraParams = make(map[string]string)
	}
	opts.ExtraParams["query"] = query
	return c.Get("/search/keyword", opts)
}

// SearchCollection 搜索合集
func (c *Client) SearchCollection(query string, opts *RequestOption) (json.RawMessage, error) {
	if opts == nil {
		opts = &RequestOption{}
	}
	if opts.ExtraParams == nil {
		opts.ExtraParams = make(map[string]string)
	}
	opts.ExtraParams["query"] = query
	return c.Get("/search/collection", opts)
}

// SearchCompany 搜索公司
func (c *Client) SearchCompany(query string, opts *RequestOption) (json.RawMessage, error) {
	if opts == nil {
		opts = &RequestOption{}
	}
	if opts.ExtraParams == nil {
		opts.ExtraParams = make(map[string]string)
	}
	opts.ExtraParams["query"] = query
	return c.Get("/search/company", opts)
}

// --- 发现 API ---

// DiscoverMovie 发现电影
func (c *Client) DiscoverMovie(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/discover/movie", opts)
}

// DiscoverTV 发现电视剧
func (c *Client) DiscoverTV(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/discover/tv", opts)
}

// --- 趋势 API ---

// GetTrending 获取趋势内容
func (c *Client) GetTrending(mediaType, timeWindow string, opts *RequestOption) (json.RawMessage, error) {
	return c.Get(fmt.Sprintf("/trending/%s/%s", mediaType, timeWindow), opts)
}

// --- 其他 API ---

// GetGenreMovieList 获取电影类型列表
func (c *Client) GetGenreMovieList(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/genre/movie/list", opts)
}

// GetGenreTVList 获取电视剧类型列表
func (c *Client) GetGenreTVList(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/genre/tv/list", opts)
}

// GetConfiguration 获取 API 配置
func (c *Client) GetConfiguration(opts *RequestOption) (json.RawMessage, error) {
	return c.Get("/configuration", opts)
}

// FindByExternalID 通过外部 ID 查找
func (c *Client) FindByExternalID(externalID, externalSource string, opts *RequestOption) (json.RawMessage, error) {
	if opts == nil {
		opts = &RequestOption{}
	}
	if opts.ExtraParams == nil {
		opts.ExtraParams = make(map[string]string)
	}
	opts.ExtraParams["external_source"] = externalSource
	return c.Get(fmt.Sprintf("/find/%s", externalID), opts)
}
