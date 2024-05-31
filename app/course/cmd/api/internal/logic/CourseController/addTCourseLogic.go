package CourseController

import (
	"context"
	"fmt"
	"github.com/moker/app/course/cmd/rpc/pb"

	"github.com/moker/app/course/cmd/api/internal/svc"
	"github.com/moker/app/course/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTCourseLogic {
	return &AddTCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTCourseLogic) AddTCourse(req *types.AddTCourseReq) (resp *types.AddTCourseResp, err error) {
	// todo: add your logic here and delete this line
	info, err := l.svcCtx.CourseCenterRPC.PutTCourse(l.ctx, &pb.TputCRequest{
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
	GetTCourseResp, err := l.svcCtx.CourseCenterRPC.GetTCourse(l.ctx, &pb.TgetCRequest{
		TId: req.Course.TeaId,
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
	return &types.AddTCourseResp{
		TCourses: courses,
	}, nil
}
