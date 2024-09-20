package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	"github.com/moker/app/usercenter/cmd/rpc/pb"
	"github.com/moker/common/ctxdata"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

//var ErrGenerateTokenError = merr.NewErrMsg("生成token失败")
//var ErrUsernamePwdError = merr.NewErrMsg("账号或密码不正确")

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	// todo: add your logic here and delete this line
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId)
	if err != nil {

	}

	return &pb.GenerateTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {

	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + seconds
	claims[ctxdata.CtxKeyJwtUserId] = userId
	fmt.Println(userId)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
