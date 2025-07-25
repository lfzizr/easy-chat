package logic

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"gitee.com/lfzizr/easy-chat/apps/user/models"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/internal/svc"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/user"
	"gitee.com/lfzizr/easy-chat/pkg/ctxdata"
	"gitee.com/lfzizr/easy-chat/pkg/encrypt"
	"gitee.com/lfzizr/easy-chat/pkg/wuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}
var (
	ErrPhoneIsRegister = errors.New("手机号已注册")
)
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line
	// 1. 校验参数
	// 2. 校验手机号是否存在
	// 3. 生成uid
	// 4. 生成token
	// 5. 保存用户信息
	// 6. 返回token
	userEntity,err := l.svcCtx.UserModel.FindByPhone(l.ctx,in.Phone,)
	if err != nil && err != models.ErrNotFound {
		return nil,err
	}
	if userEntity != nil {
		return nil,ErrPhoneIsRegister
	}
	userEntity = &models.Users{
		Id: wuid.GenUid(l.svcCtx.Config.MySQL.DataSource),
		Avatar: in.Avatar,
		Nickname: in.Nickname,
		Phone: in.Phone,
		Sex: sql.NullInt64{
			Int64: int64(in.Sex),
			Valid: true,
		},
	}

	if len(in.Password) > 0 {
		genPasswd, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, err
		}
		userEntity.Password = sql.NullString{
			String: string(genPasswd),
			Valid:  true,
		}
	}
	// 2. 保存用户信息
	_,err = l.svcCtx.UserModel.Insert(l.ctx,userEntity)
	if err != nil {
		return nil, err
	}
	// 3. 生成并返回token
	now := time.Now().Unix()
	token,err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret,now,l.svcCtx.Config.Jwt.AccessExpire,userEntity.Id)
	if err != nil {
		return nil,err
	}
	return &user.RegisterResp{
		Token: token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
