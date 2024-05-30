package UserCenter

import (
	"context"
	"github.com/moker/app/usercenter/cmd/api/internal/svc"
	"github.com/moker/app/usercenter/cmd/api/internal/types"
	"github.com/moker/app/usercenter/cmd/rpc/pb"
	"github.com/moker/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentInfoLogic {
	return &UpdateStudentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStudentInfoLogic) UpdateStudentInfo(req *types.UpdateStudentInfoReq) (resp *types.UpdateStudentInfoResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)
	info, err := l.svcCtx.UserCenterRPC.UpdateStudentInfo(l.ctx, &pb.UpdateStudentInfoReq{
		StuId:   req.Id,
		UserId:  userId,
		Name:    req.Name,
		Major:   req.Major,
		Faculty: req.Faculty,
		School:  req.School,
	})
	if err != nil {
		return nil, err
	}
	newrole := new(types.Role)
	copyStudentRoleInfo(info, newrole)
	return &types.UpdateStudentInfoResp{
		Role: *newrole,
	}, nil
}

func copyStudentRoleInfo(info *pb.UpdateStudentInfoResp, newrole *types.Role) {
	newrole.Id = info.Userinfo.Role.GetStuId()
	newrole.UserId = info.Userinfo.Role.GetUserId()
	newrole.Name = info.Userinfo.Role.GetName()
	newrole.Faculty = info.Userinfo.Role.GetFaculty()
	newrole.Major = info.Userinfo.Role.GetMajor()
	newrole.School = info.Userinfo.Role.GetSchool()
}
