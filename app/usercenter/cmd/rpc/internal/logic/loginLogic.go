package logic

import (
	"context"
	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	"github.com/moker/app/usercenter/cmd/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: 用户校验
	userData, _ := l.svcCtx.UserModel.FindOne(l.ctx, in.Account)
	if userData == nil {

	}
	saltData, err := l.svcCtx.UserSaltModel.FindOneByUserId(l.ctx, userData.Id)
	if err != nil {

	}
	isOK, err := l.svcCtx.Etools.Codedata.DeCode(in.Password, userData.Password, saltData.Salt)
	if err != nil {

	}
	if !isOK {

	}
	// todo: 凭证生成
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: userData.Id,
	})
	if err != nil {

	}
	return &pb.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
