package logic

import (
	"context"

	"github.com/moker/app/course/cmd/rpc/internal/svc"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSCourseLogic {
	return &DeleteSCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSCourseLogic) DeleteSCourse(in *pb.SdelCRequest) (*pb.SdelCReply, error) {
	// todo: add your logic here and delete this line

	return &pb.SdelCReply{}, nil
}
