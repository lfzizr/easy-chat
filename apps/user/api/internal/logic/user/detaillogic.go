package user

import (
	"context"

	"gitee.com/lfzizr/easy-chat/apps/user/api/internal/svc"
	"gitee.com/lfzizr/easy-chat/apps/user/api/internal/types"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/user"
	"gitee.com/lfzizr/easy-chat/pkg/ctxdata"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)
	getUserInfoResp, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Id: uid,
	})
	if err!= nil {
		return nil, err
	}
	/*
	var resp *types.UserInfoResp
	copier.Copy(resp, getUserInfoResp)
	*/
	return &types.UserInfoResp{
		Info: types.User{
			Id:       getUserInfoResp.User.Id,
			Mobile:   getUserInfoResp.User.Phone,
			Nickname: getUserInfoResp.User.Nickname,
			Sex:      byte(getUserInfoResp.User.Sex),
			Avatar:   getUserInfoResp.User.Avatar,
		},
	}, nil
}
