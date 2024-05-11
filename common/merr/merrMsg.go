package merr

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var message map[uint32]error

func init() {
	message = make(map[uint32]error)
	message[SERVER_COMMON_ERROR] = errors.New("服务器内部错误")
	message[DB_ERROR] = errors.New("数据库错误")
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = errors.New("数据库更新影响行数为0")
	message[REUQEST_PARAM_ERROR] = errors.New("请求参数错误")
	message[TOKEN_EXPIRE_ERROR] = errors.New("token过期")
}

func CreateNewErrorInfo(code uint32, msg string) bool {
	return NewMapErrMsg(code, msg)
}

func NewMapErrMsg(errccode uint32, errMsg string) bool {
	_, ok := message[errccode]
	if ok {
		logx.Errorf("errcode:%v is exit", errccode)
	} else {
		message[errccode] = errors.New(errMsg)
	}
	return ok
}

func StringMapErrMsg(errcode uint32) string {
	if err, ok := message[errcode]; ok {
		return err.Error()
	} else {
		return "未定义错误"
	}
}
func MapErrMsg(errcode uint32) error {
	if err, ok := message[errcode]; ok {
		return err
	} else {
		return errors.New("未定义错误")
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
