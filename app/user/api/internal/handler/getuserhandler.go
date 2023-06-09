package handler

import (
	"means/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"means/app/user/api/internal/logic"
	"means/app/user/api/internal/svc"
	"means/app/user/api/internal/types"
)

func getUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		result.HttpResult(r, w, resp, err)
	}
}
