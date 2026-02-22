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

type UpdateMovieLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMovieLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMovieLogic {
	return &UpdateMovieLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMovieLogic) UpdateMovie(req *types.AdminUpdateReq) error {
	if req.Id <= 0 {
		return errors.New("无效电影 ID")
	}

	var movie model.Movie
	if err := l.svcCtx.DB.Where("tmdb_id = ?", req.Id).First(&movie).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("电影未同步到本地，请先打开详情页")
		}
		return err
	}

	patch := make(map[string]interface{})
	updates := map[string]interface{}{
		"is_modified": true,
	}

	if req.Title != nil {
		value := trimPtrString(req.Title)
		patch["title"] = value
		updates["title"] = value
	}
	if req.OriginalTitle != nil {
		value := trimPtrString(req.OriginalTitle)
		patch["original_title"] = value
		updates["original_title"] = value
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
	if req.ReleaseDate != nil {
		value := trimPtrString(req.ReleaseDate)
		patch["release_date"] = value
		updates["release_date"] = value
	}
	if req.Status != nil {
		value := trimPtrString(req.Status)
		patch["status"] = value
		updates["status"] = value
	}
	if req.Runtime != nil {
		patch["runtime"] = *req.Runtime
		updates["runtime"] = *req.Runtime
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
	if req.GenreNames != nil {
		genres := buildGenresFromNames(req.GenreNames)
		patch["genres"] = genres
	}

	if len(patch) == 0 {
		return errors.New("没有可更新的电影字段")
	}

	mergedTMDBData, err := mergeTMDBData(movie.TmdbData, patch)
	if err != nil {
		return err
	}
	localData, err := toRawJSON(patch)
	if err != nil {
		return err
	}
	updates["tmdb_data"] = mergedTMDBData
	updates["local_data"] = localData

	return l.svcCtx.DB.Model(&model.Movie{}).Where("tmdb_id = ?", req.Id).Updates(updates).Error
}
