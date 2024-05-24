package logic

import (
	"context"
	"github.com/moker/app/usercenter/cmd/rpc/internal/svc"
	"github.com/moker/app/usercenter/cmd/rpc/pb"
	"github.com/moker/common/merr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVerCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVerCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVerCodeLogic {
	return &GetVerCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVerCodeLogic) GetVerCode(in *pb.GetVerCodeReq) (*pb.GetVerCodeResp, error) {
	// todo: add your logic here and delete this line
	code, err := l.svcCtx.Etools.VerBody.CreatVerifyCode()
	if err != nil {
		return nil, merr.Wrap(merr.VERCODE_ERROR, merr.StringMapErrMsg(merr.VERCODE_ERROR), err)
	}
	return &pb.GetVerCodeResp{
		Key:  code.ID,
		B64S: code.Base64,
	}, nil
}
