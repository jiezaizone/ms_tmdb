package tvseason

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ms_tmdb/internal/logic/tvseason"
	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"
)

func GetTvEpisodeCreditsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TvEpisodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tvseason.NewGetTvEpisodeCreditsLogic(r.Context(), svcCtx)
		err := l.GetTvEpisodeCredits(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
