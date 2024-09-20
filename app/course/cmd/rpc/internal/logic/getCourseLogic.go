package logic

import (
	"context"

	"github.com/moker/app/course/cmd/rpc/internal/svc"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseLogic {
	return &GetCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCourseLogic) GetCourse(in *pb.CourseRequest) (*pb.CourseReplyList, error) {
	// todo: add your logic here and delete this line

	return &pb.CourseReplyList{}, nil
}
