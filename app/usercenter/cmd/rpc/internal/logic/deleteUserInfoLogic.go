package logic

import (
	"context"

	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	"github.com/moker/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserInfoLogic {
	return &DeleteUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserInfoLogic) DeleteUserInfo(in *pb.DeleteUserInfoReq) (*pb.DeleteUserInfoResp, error) {
	// todo:

	return &pb.DeleteUserInfoResp{}, nil
}
