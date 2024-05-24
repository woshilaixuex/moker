package security

import (
	"github.com/moker/common/security/encrypt"
	"github.com/moker/common/security/vercode"
)

type (
	EncryptTools struct {
		Codedata *encrypt.CodeData
		VerBody  *vercode.VerBody
	}
)

func NewEncryptTools() *EncryptTools {
	return &EncryptTools{
		Codedata: encrypt.NewCodeData(),
		VerBody:  vercode.NewVerBody(),
	}
}
