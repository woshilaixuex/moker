package CourseController

import (
	"context"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/moker/app/course/cmd/api/internal/svc"
	"github.com/moker/app/course/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletSCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletSCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletSCourseLogic {
	return &DeletSCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletSCourseLogic) DeletSCourse(req *types.DeletSCourseReq) (resp *types.DeletSCourseResp, err error) {
	// todo: add your logic here and delete this line
	info, err := l.svcCtx.CourseCenterRPC.DeleteSCourse(l.ctx, &pb.SdelCRequest{
		StuId: req.StuId,
		CId:   req.CId,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeletSCourseResp{
		Info: info.GetInfo(),
	}, nil
}
