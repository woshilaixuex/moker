package vercode

import (
	"errors"
	bs64 "github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type VerCodeTool interface {
	init()
	CreatVerifyCode() (*VerBody, error)
	GetVerifyCode(code string) (string, error)
}

type VerBody struct {
	ID     string
	Base64 string
	Code   string
}

func NewVerBody() *VerBody {
	return &VerBody{}
}
func init() {
	driver := bs64.NewDriverString(Height, Width, Length, ShowLineOptions, NoiseCount, Source, &BgColor, FontsStorage, Fonts)
	captcha = bs64.NewCaptcha(driver, bs64.DefaultMemStore)
}
func (verBody *VerBody) CreatVerifyCode() (*VerBody, error) {
	var err error
	verBody.ID, verBody.Base64, verBody.Code, err = captcha.Generate()
	if err != nil {
		logx.Debug(err)
		return nil, errors.New("验证码生成失败")
	}
	//vcode := captcha.Store.Get(id, true)
	//fmt.Println(vcode)
	return verBody, nil
}
func (verBody *VerBody) GetVerifyCode(id string) (string, error) {
	if id == "" {
		return "", errors.New("id为空")
	}
	vcode := captcha.Store.Get(id, true)
	return vcode, nil
}
