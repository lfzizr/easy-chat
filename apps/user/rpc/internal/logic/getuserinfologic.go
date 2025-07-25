package logic

import (
	"context"
	"errors"

	"gitee.com/lfzizr/easy-chat/apps/user/rpc/internal/svc"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/user"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/test/model"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}
var (
	ErrUserNotFound = errors.New("不存在的用户")
)
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	userEntity, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	var resp user.UserEntity
	copier.Copy(&resp, userEntity)
	return &user.GetUserInfoResp{
		User: &resp,
	}, nil
}
