package CourseController

import (
	"context"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/moker/app/course/cmd/api/internal/svc"
	"github.com/moker/app/course/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTCourseLogic {
	return &GetTCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTCourseLogic) GetTCourse(req *types.GetTCourseReq) (resp *types.GetTCourseResp, err error) {
	// todo: add your logic here and delete this line
	GetTCourseResp, err := l.svcCtx.CourseCenterRPC.GetTCourse(l.ctx, &pb.TgetCRequest{
		TId: req.TeaId,
	})
	if err != nil {
		return nil, err
	}
	courses := make([]types.TCourse, 0)
	for _, reply := range GetTCourseResp.Replies {
		c := new(types.TCourse)
		c.Course = *new(types.Course)
		setTCourse(&c.Course, reply)
		c.Respon = reply.Respon
		courses = append(courses, *c)
	}
	return &types.GetTCourseResp{
		TCourses: courses,
	}, nil
}
func setTCourse(c *types.Course, reply *pb.TgetCReply) {
	c.CId = reply.GetCId()
	c.Info = reply.GetInfo()
	c.CName = reply.GetCName()
	c.TeaId = reply.GetTId()
	c.TeaName = reply.GetTName()
}
