package CourseController

import (
	"context"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/moker/app/course/cmd/api/internal/svc"
	"github.com/moker/app/course/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTCourseLogic {
	return &DeleteTCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTCourseLogic) DeleteTCourse(req *types.DeleteTCourseReq) (resp *types.DeleteTCourseResp, err error) {
	// todo: add your logic here and delete this line
	info, err := l.svcCtx.CourseCenterRPC.DeleteTCourse(l.ctx, &pb.TdeleteCRequest{
		TId: req.TeaId,
		CId: req.CId,
	})
	if err != nil {
		return nil, err
	}
	GetTCourseResp, err := l.svcCtx.CourseCenterRPC.GetTCourse(l.ctx, &pb.TgetCRequest{
		TId: req.TeaId,
	})
	if err != nil {
		return nil, err
	}
	courses := make([]types.TCourse, 0)
	if info.Info == "成功" {
		for _, reply := range GetTCourseResp.Replies {
			c := new(types.TCourse)
			c.Course = *new(types.Course)
			setTCourse(&c.Course, reply)
			c.Respon = reply.Respon
			courses = append(courses, *c)
		}
	}
	return &types.DeleteTCourseResp{
		TCourses: courses,
	}, nil
}
