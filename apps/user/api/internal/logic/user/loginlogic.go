package user

import (
	"context"

	"gitee.com/lfzizr/easy-chat/apps/user/api/internal/svc"
	"gitee.com/lfzizr/easy-chat/apps/user/api/internal/types"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登入
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	loginResp, err := l.svcCtx.UserClient.Login(l.ctx, &userclient.LoginReq{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Token:  loginResp.Token,
		Expire: loginResp.Expire,
	}, nil
}
