package UserController

import (
	"github.com/moker/app/usercenter/cmd/api/internal/logic/UserController"
	"github.com/moker/app/usercenter/cmd/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func VerCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := UserController.NewVerCodeLogic(r.Context(), svcCtx)
		resp, err := l.VerCode()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
