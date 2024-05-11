package result

import (
	"github.com/moker/common/merr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err *merr.MError) {
	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		causeErr := errors.Cause(err)

	}
}
