package UserController

import (
	"context"
	"github.com/moker/app/usercenter/cmd/api/internal/svc"
	"github.com/moker/app/usercenter/cmd/api/internal/types"
	"github.com/moker/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: 调用rpc获取登录服务
	loginResponse, err := l.svcCtx.UserCenterRPC.Login(l.ctx, &usercenter.LoginReq{
		AuthType: "web",
		Account:  req.Account,
		Password: req.PassWord,
	})
	if err != nil {
		logx.Error(err.Error())
		return nil, err
	}
	return &types.LoginResp{
		AccessToken:  loginResponse.AccessToken,
		AccessExpire: loginResponse.AccessExpire,
		RefreshAfter: loginResponse.RefreshAfter,
	}, nil
}
