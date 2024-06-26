package CourseController

import (
	"context"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/moker/app/course/cmd/api/internal/svc"
	"github.com/moker/app/course/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSCourseLogic {
	return &AddSCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSCourseLogic) AddSCourse(req *types.AddSCourseReq) (resp *types.AddSCourseResp, err error) {
	// todo: add your logic here and delete this line
	info, err := l.svcCtx.CourseCenterRPC.PutSCourse(l.ctx, &pb.SputCRequest{
		StuId: req.StuId,
		CId:   req.CId,
	})
	if err != nil {
		return nil, err
	}
	course, err := l.svcCtx.CourseCenterRPC.GetSearchECourse(l.ctx, &pb.SearchERequest{
		StuId: req.StuId,
	})
	if err != nil {
		return nil, err
	}
	courses := make([]types.Course, 0)
	if info.Info == "成功" {
		for _, reply := range course.Replies {
			c := new(types.Course)
			setSearchECourse(c, reply)
			courses = append(courses, *c)
		}
	}
	return &types.AddSCourseResp{
		Courses: courses,
	}, nil
}
