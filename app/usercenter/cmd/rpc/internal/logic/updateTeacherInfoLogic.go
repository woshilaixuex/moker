package logic

import (
	"context"

	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	"github.com/moker/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTeacherInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTeacherInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTeacherInfoLogic {
	return &UpdateTeacherInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTeacherInfoLogic) UpdateTeacherInfo(in *pb.UpdateTeacherInfoReq) (*pb.UpdateTeacherInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateTeacherInfoResp{}, nil
}
