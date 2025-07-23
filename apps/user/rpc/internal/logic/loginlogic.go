package logic

import (
	"context"
	"errors"
	"time"

	"gitee.com/lfzizr/easy-chat/apps/user/models"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/internal/svc"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/user"
	"gitee.com/lfzizr/easy-chat/pkg/ctxdata"
	"gitee.com/lfzizr/easy-chat/pkg/encrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}
var (
	ErrPhoneNotRegister = errors.New("手机号未注册")
	ErrPasswdNotMatch = errors.New("密码不匹配")
)
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line
	// 1. 校验手机号是否注册
	userEntity,err := l.svcCtx.UserModel.FindByPhone(l.ctx,in.Phone)
	if err != nil {
		if err == models.ErrNotFound {
			return nil,ErrPhoneNotRegister
		}
		return nil,err
	}
	// 2. 校验密码
	if !encrypt.ValidatePasswordHash(in.Password,userEntity.Password.String) {
		return nil,ErrPasswdNotMatch
	}
	// 3. 生成并返回token
	now := time.Now().Unix()
	token,err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.Secret,now,l.svcCtx.Config.Jwt.Expire,userEntity.Id)
	if err != nil {
		return nil,err
	}
	return &user.LoginResp{
		Token: token,
		Expire: now + l.svcCtx.Config.Jwt.Expire,
	}, nil
}
