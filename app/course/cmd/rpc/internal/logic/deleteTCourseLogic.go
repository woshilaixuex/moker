package logic

import (
	"context"

	"github.com/moker/app/course/cmd/rpc/internal/svc"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTCourseLogic {
	return &DeleteTCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTCourseLogic) DeleteTCourse(in *pb.TdeleteCRequest) (*pb.TdeleteCReply, error) {
	// todo: add your logic here and delete this line

	return &pb.TdeleteCReply{}, nil
}
