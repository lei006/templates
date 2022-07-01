package handler

import (
	"net/http"

	"alone-server-template/internal/logic"
	"alone-server-template/internal/svc"
	"alone-server-template/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AloneHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAloneLogic(r.Context(), ctx)
		resp, err := l.Alone(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
