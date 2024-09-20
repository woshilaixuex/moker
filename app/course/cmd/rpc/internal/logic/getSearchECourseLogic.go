package logic

import (
	"context"

	"github.com/moker/app/course/cmd/rpc/internal/svc"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSearchECourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSearchECourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSearchECourseLogic {
	return &GetSearchECourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSearchECourseLogic) GetSearchECourse(in *pb.SearchERequest) (*pb.SearchEReplyList, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchEReplyList{}, nil
}
