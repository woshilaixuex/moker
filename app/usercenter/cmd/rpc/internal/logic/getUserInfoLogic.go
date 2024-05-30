package logic

import (
	"context"
	"fmt"
	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	"github.com/moker/app/usercenter/cmd/rpc/pb"
	"github.com/moker/app/usercenter/model/role"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// todo:
	userInfo := new(pb.UserInfo)
	userInfo.Role = new(pb.Role)
	user, err := l.svcCtx.UserModel.FindOneById(l.ctx, in.UserId)
	if err != nil {
	}
	userInfo.Id = user.Id
	userInfo.Account = user.Account
	userInfo.Username = user.Username
	if user.Avatar == "student" {
		one, err := l.svcCtx.StudentsModel.FindOneByUserId(l.ctx, user.Id)
		if err != nil {
			logx.Errorf(err.Error())
		}
		copyRoleInfo(userInfo, one)
	} else {
		one, err := l.svcCtx.TeahcersModel.FindOneByUserId(l.ctx, user.Id)
		if err != nil {
			logx.Errorf(err.Error())
		}
		copyRoleInfo(userInfo, one)
	}
	fmt.Println(userInfo)
	return &pb.GetUserInfoResp{
		Userinfo: userInfo,
	}, nil
}
func copyRoleInfo(userInfo *pb.UserInfo, roleInfo any) {
	var Info interface{}
	switch trole := roleInfo.(type) {
	case *role.Students:
		Info = trole
	case *role.Teachers:
		Info = trole
	default:
		fmt.Println("Unsupported role type:")
		return
	}

	if Info != nil {
		if studentsInfo, ok := Info.(*role.Students); ok {
			userInfo.Role.StuId = studentsInfo.Id
			userInfo.Role.UserId = studentsInfo.UserId
			userInfo.Role.Faculty = studentsInfo.Faculty
			userInfo.Role.Name = studentsInfo.Name
			userInfo.Role.Major = studentsInfo.Major
			userInfo.Role.School = studentsInfo.School
		} else if teachersInfo, ok := Info.(*role.Teachers); ok {
			userInfo.Role.StuId = teachersInfo.Id
			userInfo.Role.UserId = teachersInfo.UserId
			userInfo.Role.Faculty = teachersInfo.Faculty
			userInfo.Role.Name = teachersInfo.Name
			// 注意：Teachers结构体中没有Major字段，所以这里不设置userInfo.Role.Major
			userInfo.Role.School = teachersInfo.School
		}
	}
}
