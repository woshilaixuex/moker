package merr

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var message map[uint32]error

func init() {
	message = make(map[uint32]error)
	message[UNKNOWN_ERROR] = errors.New("未知错误")
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
