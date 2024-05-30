package CourseController

import (
	"net/http"

	"github.com/moker/app/course/cmd/api/internal/logic/CourseController"
	"github.com/moker/app/course/cmd/api/internal/svc"
	"github.com/moker/app/course/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddTCourseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTCourseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := CourseController.NewAddTCourseLogic(r.Context(), svcCtx)
		resp, err := l.AddTCourse(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
