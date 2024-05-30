package CourseController

import (
	"context"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/moker/app/course/cmd/api/internal/svc"
	"github.com/moker/app/course/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSearchCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSearchCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSearchCourseLogic {
	return &GetSearchCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSearchCourseLogic) GetSearchCourse(req *types.GetSearchCourseReq) (resp *types.GetSearchCourseResp, err error) {
	// todo: add your logic here and delete this line
	course, err := l.svcCtx.CourseCenterRPC.GetSearchCourse(l.ctx, &pb.SearchACRequest{
		CName: req.CName,
	})
	if err != nil {
		return nil, err
	}
	courses := make([]types.Course, 0)
	for _, reply := range course.Replies {
		c := new(types.Course)
		setSearchCourse(c, reply)
		courses = append(courses, *c)
	}
	return &types.GetSearchCourseResp{
		Courses: courses,
	}, nil
}
func setSearchCourse(c *types.Course, reply *pb.SearchACReply) {
	c.CId = reply.GetCId()
	c.CName = reply.GetTeacherName()
	c.TeaId = reply.GetTeacherId()
	c.TeaName = reply.GetTeacherName()
	c.Info = reply.GetInfo()
}
