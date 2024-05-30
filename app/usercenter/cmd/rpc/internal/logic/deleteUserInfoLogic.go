package logic

import (
	"context"

	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	"github.com/moker/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserInfoLogic {
	return &DeleteUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserInfoLogic) DeleteUserInfo(in *pb.DeleteUserInfoReq) (*pb.DeleteUserInfoResp, error) {
	// todo:
	user, err := l.svcCtx.UserModel.FindOneById(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	if user.Avatar == "student" {
		stu, _ := l.svcCtx.StudentsModel.FindOneByUserId(l.ctx, user.Id)
		err := l.svcCtx.StudentsModel.Delete(l.ctx, stu.Id)
		if err != nil {
			return nil, err
		}
		err = l.svcCtx.UserModel.Delete(l.ctx, user.Id)
		if err != nil {
			return nil, err
		}
		err = l.svcCtx.UserSaltModel.Delete(l.ctx, user.Id)
		if err != nil {
			return nil, err
		}
	} else {
		tea, _ := l.svcCtx.TeahcersModel.FindOneByUserId(l.ctx, user.Id)
		err := l.svcCtx.TeahcersModel.Delete(l.ctx, tea.Id)
		if err != nil {
			return nil, err
		}
		err = l.svcCtx.UserModel.Delete(l.ctx, user.Id)
		if err != nil {
			return nil, err
		}
		err = l.svcCtx.UserSaltModel.Delete(l.ctx, user.Id)
		if err != nil {
			return nil, err
		}
	}
	return &pb.DeleteUserInfoResp{
		IsOK: true,
	}, nil
}
