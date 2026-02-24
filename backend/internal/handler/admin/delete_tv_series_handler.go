package admin

import (
	"net/http"

	"ms_tmdb/internal/logic/admin"
	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteTvSeriesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewDeleteTvSeriesLogic(r.Context(), svcCtx)
		if err := l.DeleteTvSeries(&req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		httpx.Ok(w)
	}
}
