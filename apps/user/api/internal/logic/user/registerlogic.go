package user

import (
	"context"

	"gitee.com/lfzizr/easy-chat/apps/user/api/internal/svc"
	"gitee.com/lfzizr/easy-chat/apps/user/api/internal/types"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	registerResp, err := l.svcCtx.UserClient.Register(l.ctx, &userclient.RegisterReq{
		Phone:    req.Phone,
		Nickname: req.Nickname,
		Password: req.Password,
		Avatar:   req.Avatar,
		Sex:      int32(req.Sex),
	})
	if err != nil {
		return nil, err
	}
	return &types.RegisterResp{
		Token:  registerResp.Token,
		Expire: registerResp.Expire,
	}, nil
}
