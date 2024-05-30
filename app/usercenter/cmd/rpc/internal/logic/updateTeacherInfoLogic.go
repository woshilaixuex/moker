package logic

import (
	"context"
	"github.com/moker/app/usercenter/model/role"

	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	"github.com/moker/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTeacherInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTeacherInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTeacherInfoLogic {
	return &UpdateTeacherInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTeacherInfoLogic) UpdateTeacherInfo(in *pb.UpdateTeacherInfoReq) (*pb.UpdateTeacherInfoResp, error) {
	// todo: add your logic here and delete this line
	tea := &role.Teachers{
		Id:      in.TeaId,
		UserId:  in.UserId,
		Name:    in.Name,
		Faculty: in.Faculty,
		School:  in.School,
	}
	err := l.svcCtx.TeahcersModel.Update(l.ctx, tea)
	if err != nil {
		return nil, err
	}
	userInfo := new(pb.UserInfo)
	userInfo.Role = new(pb.Role)
	one, err := l.svcCtx.TeahcersModel.FindOneByUserId(l.ctx, in.UserId)
	copyRoleInfo(userInfo, one)
	return &pb.UpdateTeacherInfoResp{
		Userinfo: userInfo,
	}, nil
}
