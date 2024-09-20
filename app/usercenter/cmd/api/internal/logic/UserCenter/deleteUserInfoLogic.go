package UserCenter

import (
	"context"
	"errors"
	"fmt"

	cpb "github.com/moker/app/course/cmd/rpc/pb"
	"github.com/moker/app/usercenter/cmd/rpc/pb"
	"github.com/moker/common/ctxdata"

	"github.com/moker/app/usercenter/cmd/api/internal/svc"
	"github.com/moker/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserInfoLogic {
	return &DeleteUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserInfoLogic) DeleteUserInfo(req *types.DeleteUserInfoReq) (resp *types.DeleteUserInfoResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)
	course, _ := l.svcCtx.CourseCenterRPC.GetSearchECourse(l.ctx, &cpb.SearchERequest{
		StuId: req.StuId,
	})
	idx := 0
	for _, reply := range course.Replies {
		idx++
		fmt.Println(reply)
	}
	if idx != 0 {
		return nil, errors.New("该学生有课程")
	}
	info, err := l.svcCtx.UserCenterRPC.DeleteUserInfo(l.ctx, &pb.DeleteUserInfoReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &types.DeleteUserInfoResp{
		IsOk: info.IsOK,
	}, nil
}
