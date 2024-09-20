package logic

import (
	"context"

	"github.com/moker/app/course/cmd/rpc/internal/svc"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutTCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutTCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutTCourseLogic {
	return &PutTCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutTCourseLogic) PutTCourse(in *pb.TputCRequest) (*pb.TputCReply, error) {
	// todo: add your logic here and delete this line

	return &pb.TputCReply{}, nil
}
