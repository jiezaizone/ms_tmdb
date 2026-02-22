package tvseason

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ms_tmdb/internal/logic/tvseason"
	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"
)

func GetTvSeasonVideosHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TvSeasonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tvseason.NewGetTvSeasonVideosLogic(r.Context(), svcCtx)
		err := l.GetTvSeasonVideos(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
