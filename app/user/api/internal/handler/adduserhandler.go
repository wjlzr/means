package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"means/app/user/api/internal/logic"
	"means/app/user/api/internal/svc"
	"means/app/user/api/internal/types"
)

func addUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddUserLogic(r.Context(), svcCtx)
		resp, err := l.AddUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
