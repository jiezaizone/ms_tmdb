package admin

import (
	"context"
	"errors"

	"ms_tmdb/internal/model"
	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateTvSeriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTvSeriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTvSeriesLogic {
	return &UpdateTvSeriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTvSeriesLogic) UpdateTvSeries(req *types.AdminUpdateReq) error {
	if req.Id <= 0 {
		return errors.New("无效剧集 ID")
	}

	var tv model.TVSeries
	if err := l.svcCtx.DB.Where("tmdb_id = ?", req.Id).First(&tv).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("剧集未同步到本地，请先打开详情页")
		}
		return err
	}

	patch := make(map[string]interface{})
	updates := map[string]interface{}{
		"is_modified": true,
	}

	if req.Name != nil {
		value := trimPtrString(req.Name)
		patch["name"] = value
		updates["name"] = value
	}
	if req.OriginalName != nil {
		value := trimPtrString(req.OriginalName)
		patch["original_name"] = value
		updates["original_name"] = value
	}
	if req.Overview != nil {
		value := trimPtrString(req.Overview)
		patch["overview"] = value
		updates["overview"] = value
	}
	if req.Tagline != nil {
		value := trimPtrString(req.Tagline)
		patch["tagline"] = value
		updates["tagline"] = value
	}
	if req.FirstAirDate != nil {
		value := trimPtrString(req.FirstAirDate)
		patch["first_air_date"] = value
		updates["first_air_date"] = value
	}
	if req.Status != nil {
		value := trimPtrString(req.Status)
		patch["status"] = value
		updates["status"] = value
	}
	if req.NumberOfSeasons != nil {
		patch["number_of_seasons"] = *req.NumberOfSeasons
		updates["number_of_seasons"] = *req.NumberOfSeasons
	}
	if req.NumberOfEpisodes != nil {
		patch["number_of_episodes"] = *req.NumberOfEpisodes
		updates["number_of_episodes"] = *req.NumberOfEpisodes
	}
	if req.OriginalLanguage != nil {
		value := trimPtrString(req.OriginalLanguage)
		patch["original_language"] = value
		updates["original_language"] = value
	}
	if req.Homepage != nil {
		value := trimPtrString(req.Homepage)
		patch["homepage"] = value
		updates["homepage"] = value
	}
	if req.PosterPath != nil {
		value := trimPtrString(req.PosterPath)
		patch["poster_path"] = value
		updates["poster_path"] = value
	}
	if req.BackdropPath != nil {
		value := trimPtrString(req.BackdropPath)
		patch["backdrop_path"] = value
		updates["backdrop_path"] = value
	}
	if req.VoteAverage != nil {
		patch["vote_average"] = *req.VoteAverage
		updates["vote_average"] = *req.VoteAverage
	}
	if req.Popularity != nil {
		patch["popularity"] = *req.Popularity
		updates["popularity"] = *req.Popularity
	}
	if req.Type != nil {
		value := trimPtrString(req.Type)
		patch["type"] = value
		updates["type"] = value
	}
	if req.GenreNames != nil {
		genres := buildGenresFromNames(req.GenreNames)
		patch["genres"] = genres
	}

	if len(patch) == 0 {
		return errors.New("没有可更新的剧集字段")
	}

	mergedTMDBData, err := mergeTMDBData(tv.TmdbData, patch)
	if err != nil {
		return err
	}
	localData, err := toRawJSON(patch)
	if err != nil {
		return err
	}
	updates["tmdb_data"] = mergedTMDBData
	updates["local_data"] = localData

	return l.svcCtx.DB.Model(&model.TVSeries{}).Where("tmdb_id = ?", req.Id).Updates(updates).Error
}
