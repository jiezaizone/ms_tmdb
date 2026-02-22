package person

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ms_tmdb/internal/logic/person"
	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"
)

func GetPersonCombinedCreditsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := person.NewGetPersonCombinedCreditsLogic(r.Context(), svcCtx)
		err := l.GetPersonCombinedCredits(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
