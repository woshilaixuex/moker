package CourseController

import (
	"context"
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
	GetTCourseResp, err := l.svcCtx.CourseCenterRPC.GetTCourse(l.ctx, &pb.TgetCRequest{
		TId: req.Course.TeaId,
	})
	courses := make([]types.TCourse, 0)
	if info.Info == "成功" {
		for _, reply := range GetTCourseResp.Replies {
			c := new(types.TCourse)
			c.Course = *new(types.Course)
			setTCourse(&c.Course, reply)
			c.Respon = reply.GetRespon()
			courses = append(courses, *c)
		}
	}
	return &types.UpdateTCourseResp{
		TCourses: courses,
	}, nil
}
func setUpdateTCourse(c *types.Course, reply *pb.TgetCReply) {
	c.CId = reply.GetCId()
	c.CName = reply.GetCName()
	c.TeaId = reply.GetTId()
	c.TeaName = reply.GetTName()
	c.Info = reply.GetInfo()
}
