package logic

import (
	"context"
	"time"

	"gitee.com/lfzizr/easy-chat/apps/user/models"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/internal/svc"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/user"
	"gitee.com/lfzizr/easy-chat/pkg/ctxdata"
	"gitee.com/lfzizr/easy-chat/pkg/encrypt"
	"gitee.com/lfzizr/easy-chat/pkg/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}
var (
	ErrPhoneNotRegister = xerr.New(xerr.SERVER_COMMON_ERR,"手机号未注册")
	ErrPasswdNotMatch = xerr.New(xerr.SERVER_COMMON_ERR,"密码不匹配")
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
			return nil,errors.Wrap(ErrPhoneNotRegister,ErrPhoneNotRegister.Error())
		}
		return nil,errors.Wrapf(xerr.NewDBErr(),"find user by phone err %v ,req %v",err,in.Phone)
	}
	// 2. 校验密码
	if !encrypt.ValidatePasswordHash(in.Password,userEntity.Password.String) {
		return nil,errors.Wrap(ErrPasswdNotMatch,ErrPasswdNotMatch.Error())
	}
	// 3. 生成并返回token
	now := time.Now().Unix()
	token,err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret,now,l.svcCtx.Config.Jwt.AccessExpire,userEntity.Id)
	if err != nil {
		return nil,errors.Wrapf(xerr.NewInternalErr(),"ctxdata get jwt err: %v",err)
	}
	return &user.LoginResp{
		Token: token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
