package logic

import (
	"context"

	"github.com/moker/app/course/cmd/rpc/internal/svc"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutSCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutSCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutSCourseLogic {
	return &PutSCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutSCourseLogic) PutSCourse(in *pb.SputCRequest) (*pb.SputCReply, error) {
	// todo: add your logic here and delete this line

	return &pb.SputCReply{}, nil
}
