package UserCenter

import (
	"context"
	"github.com/moker/app/usercenter/cmd/rpc/pb"
	"github.com/moker/common/ctxdata"

	"github.com/moker/app/usercenter/cmd/api/internal/svc"
	"github.com/moker/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)
	info, err := l.svcCtx.UserCenterRPC.GetUserInfo(l.ctx, &pb.GetUserInfoReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &types.GetUserInfoResp{
		Username: info.Userinfo.Username,
		Role: types.Role{
			Id:      info.Userinfo.Role.StuId,
			UserId:  info.Userinfo.Role.UserId,
			Name:    info.Userinfo.Role.Name,
			Major:   info.Userinfo.Role.Major,
			Faculty: info.Userinfo.Role.Faculty,
			School:  info.Userinfo.Role.School,
		},
	}, nil
}
