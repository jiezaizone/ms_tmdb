package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"ms_tmdb/internal/logic/proxy"
	"ms_tmdb/pkg/tmdbclient"
)

type tmdbExactHandler func(opts *tmdbclient.RequestOption, r *http.Request) (json.RawMessage, error)

type tmdbPatternHandler func(matches []string, opts *tmdbclient.RequestOption, r *http.Request) (json.RawMessage, error)

type tmdbPatternRoute struct {
	pattern *regexp.Regexp
	handler tmdbPatternHandler
}

type tmdbRouteDispatcher struct {
	client   *tmdbclient.Client
	exacts   map[string]tmdbExactHandler
	patterns []tmdbPatternRoute
}

func newTmdbRouteDispatcher(client *tmdbclient.Client, proxyService *proxy.ProxyService) *tmdbRouteDispatcher {
	return &tmdbRouteDispatcher{
		client: client,
		exacts: buildExactHandlers(client),
		patterns: []tmdbPatternRoute{
			{pattern: regexp.MustCompile(`^/movie/(-?\d+)$`), handler: detailRoute(proxyService.GetMovieDetail)},
			{pattern: regexp.MustCompile(`^/movie/(-?\d+)/(.+)$`), handler: passThroughRoute(client, "/movie/%d/%s", 1)},
			{pattern: regexp.MustCompile(`^/tv/(-?\d+)/season/(\d+)/episode/(\d+)/(.+)$`), handler: passThroughRoute(client, "/tv/%d/season/%d/episode/%d/%s", 1, 2, 3)},
			{pattern: regexp.MustCompile(`^/tv/(-?\d+)/season/(\d+)/episode/(\d+)$`), handler: episodeDetailRoute(client.GetTVEpisode)},
			{pattern: regexp.MustCompile(`^/tv/(-?\d+)/season/(\d+)/(.+)$`), handler: passThroughRoute(client, "/tv/%d/season/%d/%s", 1, 2)},
			{pattern: regexp.MustCompile(`^/tv/(-?\d+)/season/(\d+)$`), handler: seasonDetailRoute(proxyService.GetTvSeasonDetail)},
			{pattern: regexp.MustCompile(`^/tv/(-?\d+)$`), handler: detailRoute(proxyService.GetTvSeriesDetail)},
			{pattern: regexp.MustCompile(`^/tv/(-?\d+)/(.+)$`), handler: passThroughRoute(client, "/tv/%d/%s", 1)},
			{pattern: regexp.MustCompile(`^/person/(-?\d+)$`), handler: detailRoute(proxyService.GetPersonDetail)},
			{pattern: regexp.MustCompile(`^/person/(-?\d+)/(.+)$`), handler: passThroughRoute(client, "/person/%d/%s", 1)},
			{pattern: regexp.MustCompile(`^/trending/(\w+)/(\w+)$`), handler: trendingRoute(client.GetTrending)},
			{pattern: regexp.MustCompile(`^/find/(.+)$`), handler: findRoute(client.FindByExternalID)},
		},
	}
}

func (d *tmdbRouteDispatcher) dispatch(path string, opts *tmdbclient.RequestOption, r *http.Request) (json.RawMessage, error) {
	if handler, ok := d.exacts[path]; ok {
		return handler(opts, r)
	}

	for _, route := range d.patterns {
		if matches := route.pattern.FindStringSubmatch(path); matches != nil {
			return route.handler(matches, opts, r)
		}
	}

	return passThroughPath(d.client, path, opts)
}

func buildExactHandlers(client *tmdbclient.Client) map[string]tmdbExactHandler {
	return map[string]tmdbExactHandler{
		"/search/movie":      searchRoute(client.SearchMovie),
		"/search/tv":         searchRoute(client.SearchTV),
		"/search/person":     searchRoute(client.SearchPerson),
		"/search/multi":      searchRoute(client.SearchMulti),
		"/search/keyword":    searchRoute(client.SearchKeyword),
		"/search/collection": searchRoute(client.SearchCollection),
		"/search/company":    searchRoute(client.SearchCompany),
		"/discover/movie":    optionOnlyRoute(client.DiscoverMovie),
		"/discover/tv":       optionOnlyRoute(client.DiscoverTV),
		"/genre/movie/list":  optionOnlyRoute(client.GetGenreMovieList),
		"/genre/tv/list":     optionOnlyRoute(client.GetGenreTVList),
		"/movie/now_playing": optionOnlyRoute(client.GetNowPlayingMovies),
		"/movie/popular":     optionOnlyRoute(client.GetPopularMovies),
		"/movie/top_rated":   optionOnlyRoute(client.GetTopRatedMovies),
		"/movie/upcoming":    optionOnlyRoute(client.GetUpcomingMovies),
		"/tv/airing_today":   optionOnlyRoute(client.GetAiringTodayTV),
		"/tv/on_the_air":     optionOnlyRoute(client.GetOnTheAirTV),
		"/tv/popular":        optionOnlyRoute(client.GetPopularTV),
		"/tv/top_rated":      optionOnlyRoute(client.GetTopRatedTV),
		"/person/popular":    optionOnlyRoute(client.GetPopularPeople),
		"/configuration":     optionOnlyRoute(client.GetConfiguration),
	}
}

func optionOnlyRoute(handler func(opts *tmdbclient.RequestOption) (json.RawMessage, error)) tmdbExactHandler {
	return func(opts *tmdbclient.RequestOption, _ *http.Request) (json.RawMessage, error) {
		return handler(opts)
	}
}

func searchRoute(handler func(query string, opts *tmdbclient.RequestOption) (json.RawMessage, error)) tmdbExactHandler {
	return func(opts *tmdbclient.RequestOption, r *http.Request) (json.RawMessage, error) {
		return handler(r.URL.Query().Get("query"), opts)
	}
}

func detailRoute(handler func(id int, opts *tmdbclient.RequestOption) (json.RawMessage, error)) tmdbPatternHandler {
	return func(matches []string, opts *tmdbclient.RequestOption, _ *http.Request) (json.RawMessage, error) {
		id, err := parseIntParam(matches[1], "id")
		if err != nil {
			return nil, err
		}
		return handler(id, opts)
	}
}

func seasonDetailRoute(handler func(seriesID, seasonNum int, opts *tmdbclient.RequestOption) (json.RawMessage, error)) tmdbPatternHandler {
	return func(matches []string, opts *tmdbclient.RequestOption, _ *http.Request) (json.RawMessage, error) {
		seriesID, err := parseIntParam(matches[1], "series_id")
		if err != nil {
			return nil, err
		}
		seasonNum, err := parseIntParam(matches[2], "season_number")
		if err != nil {
			return nil, err
		}
		return handler(seriesID, seasonNum, opts)
	}
}

func episodeDetailRoute(handler func(seriesID, seasonNum, episodeNum int, opts *tmdbclient.RequestOption) (json.RawMessage, error)) tmdbPatternHandler {
	return func(matches []string, opts *tmdbclient.RequestOption, _ *http.Request) (json.RawMessage, error) {
		seriesID, err := parseIntParam(matches[1], "series_id")
		if err != nil {
			return nil, err
		}
		seasonNum, err := parseIntParam(matches[2], "season_number")
		if err != nil {
			return nil, err
		}
		episodeNum, err := parseIntParam(matches[3], "episode_number")
		if err != nil {
			return nil, err
		}
		return handler(seriesID, seasonNum, episodeNum, opts)
	}
}

func passThroughRoute(client *tmdbclient.Client, pattern string, intParamIndexes ...int) tmdbPatternHandler {
	indexes := map[int]bool{}
	for _, index := range intParamIndexes {
		indexes[index] = true
	}

	return func(matches []string, opts *tmdbclient.RequestOption, _ *http.Request) (json.RawMessage, error) {
		values := make([]any, 0, len(matches)-1)
		for i := 1; i < len(matches); i++ {
			item := matches[i]
			if indexes[i] {
				intValue, err := parseIntParam(item, fmt.Sprintf("match_%d", i))
				if err != nil {
					return nil, err
				}
				values = append(values, intValue)
				continue
			}
			values = append(values, item)
		}
		return passThroughPath(client, fmt.Sprintf(pattern, values...), opts)
	}
}

func trendingRoute(handler func(mediaType, timeWindow string, opts *tmdbclient.RequestOption) (json.RawMessage, error)) tmdbPatternHandler {
	return func(matches []string, opts *tmdbclient.RequestOption, _ *http.Request) (json.RawMessage, error) {
		return handler(matches[1], matches[2], opts)
	}
}

func findRoute(handler func(externalID, externalSource string, opts *tmdbclient.RequestOption) (json.RawMessage, error)) tmdbPatternHandler {
	return func(matches []string, opts *tmdbclient.RequestOption, r *http.Request) (json.RawMessage, error) {
		externalSource := r.URL.Query().Get("external_source")
		return handler(matches[1], externalSource, opts)
	}
}

func passThroughPath(client *tmdbclient.Client, path string, opts *tmdbclient.RequestOption) (json.RawMessage, error) {
	return client.Request(path, opts)
}
