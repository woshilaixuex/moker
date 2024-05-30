package logic

import (
	"context"
	"github.com/moker/app/usercenter/model/role"

	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	"github.com/moker/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentInfoLogic {
	return &UpdateStudentInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStudentInfoLogic) UpdateStudentInfo(in *pb.UpdateStudentInfoReq) (*pb.UpdateStudentInfoResp, error) {
	// todo:
	stu := &role.Students{
		Id:      in.StuId,
		UserId:  in.UserId,
		Name:    in.Name,
		Major:   in.Major,
		Faculty: in.Faculty,
		School:  in.School,
	}

	err := l.svcCtx.StudentsModel.Update(l.ctx, stu)
	if err != nil {
		return nil, err
	}
	userInfo := new(pb.UserInfo)
	userInfo.Role = new(pb.Role)
	one, err := l.svcCtx.StudentsModel.FindOneByUserId(l.ctx, in.UserId)
	copyRoleInfo(userInfo, one)
	return &pb.UpdateStudentInfoResp{
		Userinfo: userInfo,
	}, nil
}
