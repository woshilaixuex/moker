package merr

import (
	"fmt"
	"time"
)

// MError 重写后的error,方便时间收集上报
type MError struct {
	Code uint32
	When time.Time
	error
}

func (e *MError) Error() string {
	return fmt.Sprintf("%s: %s", e.When, e.Error())
}

func NewMError(code uint32) *MError {
	return &MError{
		Code:  code,
		When:  time.Now(),
		error: MapErrMsg(code),
	}
}

func (e *MError) GetErrCode() uint32 {
	return e.Code
}

func (e *MError) GetErrMsg() string {
	return e.Error()
}

func (e *MError) GetTime() time.Time {
	return e.When
}
