package logic

import (
	"context"

	"github.com/moker/app/course/cmd/rpc/internal/svc"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSearchCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSearchCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSearchCourseLogic {
	return &GetSearchCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSearchCourseLogic) GetSearchCourse(in *pb.SearchACRequest) (*pb.SearchACReplyList, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchACReplyList{}, nil
}
