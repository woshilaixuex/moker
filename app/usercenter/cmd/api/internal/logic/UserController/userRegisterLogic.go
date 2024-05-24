package UserController

import (
	"context"
	"github.com/moker/app/usercenter/cmd/api/internal/svc"
	"github.com/moker/app/usercenter/cmd/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
