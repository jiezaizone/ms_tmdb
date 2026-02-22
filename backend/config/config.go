package config

import "github.com/zeromicro/go-zero/rest"

// Config 服务配置
type Config struct {
	rest.RestConf

	// PostgreSQL 配置
	Postgres struct {
		DataSource string
	}

	// TMDB 配置
	Tmdb struct {
		ApiKey          string
		BaseURL         string
		DefaultLanguage string
		RateLimit       int
	}

	// 缓存配置
	Cache struct {
		MovieTTL  int // 电影缓存时长(小时)
		TVTTL     int // 电视剧缓存时长(小时)
		PersonTTL int // 人物缓存时长(小时)
	}
}
