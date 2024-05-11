package merr

import (
	"fmt"
	"time"
)

type BasicErr interface {
	GetErrCode() uint32
	GetErrMsg() string
	GetTime() time.Time
	error
}

// Causer 底层错误信息
type Causer struct {
	Code    uint32
	When    time.Time
	cause   error
	Message string
}

func (e *Causer) Error() string {
	return fmt.Sprintf("%s: %s", e.When, e.Message)
}

func NewCauser(code uint32, message string) BasicErr {
	return &Causer{
		Code:    code,
		When:    time.Now(),
		cause:   nil,
		Message: message,
	}
}

func (e *Causer) GetErrCode() uint32 {
	return e.Code
}

func (e *Causer) GetErrMsg() string {
	return e.Error()
}

func (e *Causer) GetTime() time.Time {
	return e.When
}

// MError 重写后的error,方便时间收集上报
type MError struct {
	Code    uint32
	When    time.Time
	cause   error
	Message string
}

func (e *MError) Unwrap() error {
	return e.cause
}

func (e *MError) Error() string {
	return fmt.Sprintf("%s: %s", e.When, e.Message)
}

func NewMError(code uint32, message string, cause error) BasicErr {
	return &MError{
		Code:    code,
		When:    time.Now(),
		cause:   cause,
		Message: message,
	}
}
func (e *MError) Cause() error {
	if e.cause == nil {
		return nil
	}
	return e.cause
}
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
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
