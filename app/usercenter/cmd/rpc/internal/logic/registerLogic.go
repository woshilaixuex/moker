package logic

import (
	"context"
	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	"github.com/moker/app/usercenter/cmd/rpc/pb"
	"github.com/moker/app/usercenter/model/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	var userId int64
	// todo: 验证码
	if in.Code == "" {

	}
	// todo: 数据校验
	userData, _ := l.svcCtx.UserModel.FindOne(l.ctx, in.Account)
	if userData != nil {

	}
	// todo: 注册事务：密码加盐+用户存储
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		userinfo := new(user.User)
		userinfo.Account = in.Account
		userinfo.Username = in.Account
		encryptData, err := l.svcCtx.Etools.Codedata.Encrypt(in.Password)
		if err != nil {

		}
		userinfo.Password = encryptData.HashCode
		insertUserResult, err := l.svcCtx.UserModel.Insert(ctx, session, userinfo)
		if err != nil {

		}
		lastId, err := insertUserResult.LastInsertId()
		if err != nil {

		}
		userId = lastId
		_, err = l.svcCtx.UserSaltModel.Insert(ctx, session, &user.UserSalt{
			UserId: lastId,
			Salt:   encryptData.Salt,
		})
		if err != nil {

		}
		return nil
	}); err != nil {

	}
	// todo: 凭证生成
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {

	}
	return &pb.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
