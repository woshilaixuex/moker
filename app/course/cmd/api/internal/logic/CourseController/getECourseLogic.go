package CourseController

import (
	"context"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/moker/app/course/cmd/api/internal/svc"
	"github.com/moker/app/course/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetECourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetECourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetECourseLogic {
	return &GetECourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetECourseLogic) GetECourse(req *types.GetECourseReq) (resp *types.GetECourseResp, err error) {
	// todo: add your logic here and delete this line
	course, err := l.svcCtx.CourseCenterRPC.GetSearchECourse(l.ctx, &pb.SearchERequest{
		StuId: req.StuId,
	})
	if err != nil {
		return nil, err
	}
	courses := make([]types.Course, 0)
	for _, reply := range course.Replies {
		c := new(types.Course)
		setSearchECourse(c, reply)
		courses = append(courses, *c)
	}
	return &types.GetECourseResp{
		Courses: courses,
	}, nil
}
func setSearchECourse(c *types.Course, reply *pb.SearchEReply) {
	c.CId = reply.GetCId()
	c.CName = reply.GetTeacherName()
	c.TeaId = reply.GetTeacherId()
	c.TeaName = reply.GetTeacherName()
	c.Info = reply.GetInfo()
}
