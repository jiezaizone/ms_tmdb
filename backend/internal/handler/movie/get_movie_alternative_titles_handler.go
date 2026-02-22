package movie

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ms_tmdb/internal/logic/movie"
	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"
)

func GetMovieAlternativeTitlesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := movie.NewGetMovieAlternativeTitlesLogic(r.Context(), svcCtx)
		err := l.GetMovieAlternativeTitles(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
