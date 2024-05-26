package UserController

import (
	"context"
	"github.com/moker/app/usercenter/cmd/api/internal/svc"
	"github.com/moker/app/usercenter/cmd/api/internal/types"
	"github.com/moker/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerCodeLogic {
	return &VerCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerCodeLogic) VerCode() (resp *types.VerCodeResp, err error) {
	// todo: 调用rpc获取验证码服务
	verResp, err := l.svcCtx.UserCenterRPC.GetVerCode(l.ctx, &usercenter.GetVerCodeReq{})
	if err != nil {
		logx.Error(err.Error())
		return nil, err
	}

	return &types.VerCodeResp{
		Base64: verResp.B64S,
		Key:    verResp.Key,
	}, nil
}
