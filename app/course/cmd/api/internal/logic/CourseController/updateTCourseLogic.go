package CourseController

import (
	"context"
	"fmt"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/moker/app/course/cmd/api/internal/svc"
	"github.com/moker/app/course/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTCourseLogic {
	return &UpdateTCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTCourseLogic) UpdateTCourse(req *types.UpdateTCourseReq) (resp *types.UpdateTCourseResp, err error) {
	// todo: add your logic here and delete this line
	info, err := l.svcCtx.CourseCenterRPC.UpdateTCourse(l.ctx, &pb.TupdateCRequest{
		TId:   req.Course.TeaId,
		TName: req.Course.TeaName,
		CId:   req.Course.CId,
		CName: req.Course.CName,
		Info:  req.Course.Info,
	})
	fmt.Println(info)
	if err != nil {
		return nil, err
	}
	return &types.UpdateTCourseResp{}, nil
}
