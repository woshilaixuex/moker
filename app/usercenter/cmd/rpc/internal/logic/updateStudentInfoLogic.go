package logic

import (
	"context"

	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	"github.com/moker/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentInfoLogic {
	return &UpdateStudentInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStudentInfoLogic) UpdateStudentInfo(in *pb.UpdateStudentInfoReq) (*pb.UpdateStudentInfoResp, error) {
	// todo:

	return &pb.UpdateStudentInfoResp{}, nil
}
