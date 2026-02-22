package proxy

import (
	"encoding/json"
	"time"

	"ms_tmdb/internal/model"
	"ms_tmdb/pkg/tmdbclient"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// ProxyService 代理服务，封装 Read-Through 缓存逻辑
type ProxyService struct {
	DB         *gorm.DB
	TmdbClient *tmdbclient.Client
}

func NewProxyService(db *gorm.DB, client *tmdbclient.Client) *ProxyService {
	return &ProxyService{DB: db, TmdbClient: client}
}

// GetMovieDetail Read-Through 获取电影详情
func (s *ProxyService) GetMovieDetail(tmdbID int, opts *tmdbclient.RequestOption) (json.RawMessage, error) {
	var movie model.Movie
	err := s.DB.Where("tmdb_id = ?", tmdbID).First(&movie).Error

	if err == nil {
		if movie.IsModified || !isExpired(movie.LastSyncedAt, 24*time.Hour) {
			return json.RawMessage(movie.TmdbData), nil
		}
	}

	data, fetchErr := s.TmdbClient.GetMovie(tmdbID, opts)
	if fetchErr != nil {
		if err == nil {
			logx.Infof("TMDB 不可用，返回本地缓存: movie/%d", tmdbID)
			return json.RawMessage(movie.TmdbData), nil
		}
		return nil, fetchErr
	}

	s.upsertMovie(tmdbID, data)
	return data, nil
}

// GetTvSeriesDetail Read-Through 获取电视剧详情
func (s *ProxyService) GetTvSeriesDetail(tmdbID int, opts *tmdbclient.RequestOption) (json.RawMessage, error) {
	var tv model.TVSeries
	err := s.DB.Where("tmdb_id = ?", tmdbID).First(&tv).Error

	if err == nil {
		if tv.IsModified || !isExpired(tv.LastSyncedAt, 24*time.Hour) {
			return json.RawMessage(tv.TmdbData), nil
		}
	}

	data, fetchErr := s.TmdbClient.GetTVSeries(tmdbID, opts)
	if fetchErr != nil {
		if err == nil {
			return json.RawMessage(tv.TmdbData), nil
		}
		return nil, fetchErr
	}

	s.upsertTVSeries(tmdbID, data)
	return data, nil
}

// GetPersonDetail Read-Through 获取人物详情
func (s *ProxyService) GetPersonDetail(tmdbID int, opts *tmdbclient.RequestOption) (json.RawMessage, error) {
	var person model.Person
	err := s.DB.Where("tmdb_id = ?", tmdbID).First(&person).Error

	if err == nil {
		if person.IsModified || !isExpired(person.LastSyncedAt, 48*time.Hour) {
			return json.RawMessage(person.TmdbData), nil
		}
	}

	data, fetchErr := s.TmdbClient.GetPerson(tmdbID, opts)
	if fetchErr != nil {
		if err == nil {
			return json.RawMessage(person.TmdbData), nil
		}
		return nil, fetchErr
	}

	s.upsertPerson(tmdbID, data)
	return data, nil
}

func (s *ProxyService) upsertMovie(tmdbID int, data json.RawMessage) {
	var parsed struct {
		Title            string  `json:"title"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		Popularity       float64 `json:"popularity"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int     `json:"vote_count"`
		PosterPath       string  `json:"poster_path"`
		BackdropPath     string  `json:"backdrop_path"`
		OriginalLanguage string  `json:"original_language"`
		Adult            bool    `json:"adult"`
		Status           string  `json:"status"`
		Runtime          int     `json:"runtime"`
		Budget           int64   `json:"budget"`
		Revenue          int64   `json:"revenue"`
		Tagline          string  `json:"tagline"`
		Homepage         string  `json:"homepage"`
		ImdbID           string  `json:"imdb_id"`
	}
	json.Unmarshal(data, &parsed)

	now := time.Now()
	result := s.DB.Where("tmdb_id = ?", tmdbID).First(&model.Movie{})
	if result.Error == gorm.ErrRecordNotFound {
		s.DB.Create(&model.Movie{
			TmdbID: tmdbID, Title: parsed.Title, OriginalTitle: parsed.OriginalTitle,
			Overview: parsed.Overview, ReleaseDate: parsed.ReleaseDate,
			Popularity: parsed.Popularity, VoteAverage: parsed.VoteAverage, VoteCount: parsed.VoteCount,
			PosterPath: parsed.PosterPath, BackdropPath: parsed.BackdropPath,
			OriginalLanguage: parsed.OriginalLanguage, Adult: parsed.Adult, Status: parsed.Status,
			Runtime: parsed.Runtime, Budget: parsed.Budget, Revenue: parsed.Revenue,
			Tagline: parsed.Tagline, Homepage: parsed.Homepage, ImdbID: parsed.ImdbID,
			TmdbData: model.RawJSON(data), LastSyncedAt: &now,
		})
	} else {
		s.DB.Model(&model.Movie{}).Where("tmdb_id = ?", tmdbID).Updates(map[string]interface{}{
			"title": parsed.Title, "original_title": parsed.OriginalTitle,
			"overview": parsed.Overview, "popularity": parsed.Popularity,
			"vote_average": parsed.VoteAverage, "poster_path": parsed.PosterPath,
			"tmdb_data": model.RawJSON(data), "last_synced_at": &now,
		})
	}
}

func (s *ProxyService) upsertTVSeries(tmdbID int, data json.RawMessage) {
	var parsed struct {
		Name         string  `json:"name"`
		OriginalName string  `json:"original_name"`
		Overview     string  `json:"overview"`
		FirstAirDate string  `json:"first_air_date"`
		Popularity   float64 `json:"popularity"`
		VoteAverage  float64 `json:"vote_average"`
		PosterPath   string  `json:"poster_path"`
		Status       string  `json:"status"`
	}
	json.Unmarshal(data, &parsed)

	now := time.Now()
	result := s.DB.Where("tmdb_id = ?", tmdbID).First(&model.TVSeries{})
	if result.Error == gorm.ErrRecordNotFound {
		s.DB.Create(&model.TVSeries{
			TmdbID: tmdbID, Name: parsed.Name, OriginalName: parsed.OriginalName,
			Overview: parsed.Overview, FirstAirDate: parsed.FirstAirDate,
			Popularity: parsed.Popularity, VoteAverage: parsed.VoteAverage,
			PosterPath: parsed.PosterPath, Status: parsed.Status,
			TmdbData: model.RawJSON(data), LastSyncedAt: &now,
		})
	} else {
		s.DB.Model(&model.TVSeries{}).Where("tmdb_id = ?", tmdbID).Updates(map[string]interface{}{
			"name": parsed.Name, "overview": parsed.Overview,
			"popularity": parsed.Popularity, "vote_average": parsed.VoteAverage,
			"poster_path": parsed.PosterPath, "tmdb_data": model.RawJSON(data), "last_synced_at": &now,
		})
	}
}

func (s *ProxyService) upsertPerson(tmdbID int, data json.RawMessage) {
	var parsed struct {
		Name        string  `json:"name"`
		Biography   string  `json:"biography"`
		Popularity  float64 `json:"popularity"`
		ProfilePath string  `json:"profile_path"`
	}
	json.Unmarshal(data, &parsed)

	now := time.Now()
	result := s.DB.Where("tmdb_id = ?", tmdbID).First(&model.Person{})
	if result.Error == gorm.ErrRecordNotFound {
		s.DB.Create(&model.Person{
			TmdbID: tmdbID, Name: parsed.Name, Biography: parsed.Biography,
			Popularity: parsed.Popularity, ProfilePath: parsed.ProfilePath,
			TmdbData: model.RawJSON(data), LastSyncedAt: &now,
		})
	} else {
		s.DB.Model(&model.Person{}).Where("tmdb_id = ?", tmdbID).Updates(map[string]interface{}{
			"name": parsed.Name, "popularity": parsed.Popularity,
			"profile_path": parsed.ProfilePath, "tmdb_data": model.RawJSON(data), "last_synced_at": &now,
		})
	}
}

func isExpired(syncedAt *time.Time, ttl time.Duration) bool {
	if syncedAt == nil {
		return true
	}
	return time.Since(*syncedAt) > ttl
}
