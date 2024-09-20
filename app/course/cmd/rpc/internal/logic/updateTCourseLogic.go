package logic

import (
	"context"

	"github.com/moker/app/course/cmd/rpc/internal/svc"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTCourseLogic {
	return &UpdateTCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTCourseLogic) UpdateTCourse(in *pb.TupdateCRequest) (*pb.TupdateCReply, error) {
	// todo: add your logic here and delete this line

	return &pb.TupdateCReply{}, nil
}
