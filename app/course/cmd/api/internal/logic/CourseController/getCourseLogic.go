package CourseController

import (
	"context"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/moker/app/course/cmd/api/internal/svc"
	"github.com/moker/app/course/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseLogic {
	return &GetCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCourseLogic) GetCourse(req *types.GetCoursReq) (resp *types.GetCoursResp, err error) {
	// todo: add your logic here and delete this line
	course, err := l.svcCtx.CourseCenterRPC.GetCourse(l.ctx, &pb.CourseRequest{})
	if err != nil {
		return nil, err
	}
	courses := make([]types.Course, 0)
	for _, reply := range course.Replies {
		c := new(types.Course)
		setCourse(c, reply)
		courses = append(courses, *c)
	}
	return &types.GetCoursResp{
		Courses: courses,
	}, nil
}
func setCourse(c *types.Course, reply *pb.CourseReply) {
	c.CId = reply.GetCId()
	c.CName = reply.GetTeacherName()
	c.TeaId = reply.GetTeacherId()
	c.TeaName = reply.GetTeacherName()
	c.Info = reply.GetInfo()
}
