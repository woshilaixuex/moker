package logic

import (
	"context"

	"github.com/moker/app/course/cmd/rpc/internal/svc"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTCourseLogic {
	return &GetTCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTCourseLogic) GetTCourse(in *pb.TgetCRequest) (*pb.TgetCReplyList, error) {
	// todo: add your logic here and delete this line

	return &pb.TgetCReplyList{}, nil
}
