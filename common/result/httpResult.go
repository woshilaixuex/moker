package result

import (
	"github.com/moker/common/merr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

// HttpResult 自定义返回体
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err *merr.MError) {
	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		errcode := merr.UNKNOWN_ERROR
		// 实现类errors库的拓展
		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*merr.Causer); ok {
			errcode = e.GetErrCode()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if merr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
				}
			}
		}
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v: %+v", err.Error(), err.GetTime())
		httpx.WriteJson(w, http.StatusBadRequest, ErrorResp(errcode))
	}
}
