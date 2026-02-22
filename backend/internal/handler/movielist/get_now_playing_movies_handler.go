package movielist

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ms_tmdb/internal/logic/movielist"
	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"
)

func GetNowPlayingMoviesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := movielist.NewGetNowPlayingMoviesLogic(r.Context(), svcCtx)
		err := l.GetNowPlayingMovies(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
