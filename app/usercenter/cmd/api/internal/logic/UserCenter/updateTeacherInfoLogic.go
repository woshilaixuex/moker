package UserCenter

import (
	"context"
	"github.com/moker/app/usercenter/cmd/rpc/pb"
	"github.com/moker/common/ctxdata"

	"github.com/moker/app/usercenter/cmd/api/internal/svc"
	"github.com/moker/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTeacherInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTeacherInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTeacherInfoLogic {
	return &UpdateTeacherInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTeacherInfoLogic) UpdateTeacherInfo(req *types.UpdateTeacherInfoReq) (resp *types.UpdateTeacherInfoResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)
	info, err := l.svcCtx.UserCenterRPC.UpdateTeacherInfo(l.ctx, &pb.UpdateTeacherInfoReq{
		TeaId:   req.Id,
		UserId:  userId,
		Name:    req.Name,
		Faculty: req.Faculty,
		School:  req.School,
	})
	if err != nil {
		return nil, err
	}
	newrole := new(types.Role)
	copyTeacherRoleInfo(info, newrole)
	return &types.UpdateTeacherInfoResp{
		Role: *newrole,
	}, nil
}
func copyTeacherRoleInfo(info *pb.UpdateTeacherInfoResp, newrole *types.Role) {
	newrole.Id = info.Userinfo.Role.GetStuId()
	newrole.UserId = info.Userinfo.Role.GetUserId()
	newrole.Name = info.Userinfo.Role.GetName()
	newrole.Faculty = info.Userinfo.Role.GetFaculty()
	newrole.Major = info.Userinfo.Role.GetMajor()
	newrole.School = info.Userinfo.Role.GetSchool()
}
