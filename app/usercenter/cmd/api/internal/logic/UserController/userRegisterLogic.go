package UserController

import (
	"context"
	"github.com/moker/app/usercenter/cmd/api/internal/svc"
	"github.com/moker/app/usercenter/cmd/api/internal/types"
	"github.com/moker/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: 调用rpc获取注册服务
	registerResponse, err := l.svcCtx.UserCenterRPC.Register(l.ctx, &usercenter.RegisterReq{
		Code:     req.VerCode,
		Key:      req.VerKey,
		Account:  req.Account,
		Password: req.PassWord,
		Role:     req.Role,
		AuthKey:  "nil",
		AuthType: "web",
	})
	if err != nil {
		logx.Error(err.Error())
		return nil, err
	}
	return &types.RegisterResp{
		AccessToken:  registerResponse.AccessToken,
		AccessExpire: registerResponse.AccessExpire,
		RefreshAfter: registerResponse.RefreshAfter,
		Role:         req.Role,
	}, nil
}
