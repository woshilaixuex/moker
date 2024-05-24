package security

import (
	"testing"
)

func TestNewEncryptTools(t *testing.T) {
	tools := NewEncryptTools()
	code, err := tools.VerBody.CreatVerifyCode()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(code.GetVerifyCode(code.ID))
}
