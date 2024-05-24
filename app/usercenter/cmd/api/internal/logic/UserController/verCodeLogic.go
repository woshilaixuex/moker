package UserController

import (
	"context"
	"github.com/moker/app/usercenter/cmd/api/internal/svc"
	"github.com/moker/app/usercenter/cmd/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
