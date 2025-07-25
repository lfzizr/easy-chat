package logic

import (
	"context"
	"fmt"

	"gitee.com/lfzizr/easy-chat/apps/user/models"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/internal/svc"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)
type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	// todo: add your logic here and delete this line
	var (
		userEntities []*models.Users
		err error
	)
	if in.Phone != "" {
		userEntity, err := l.svcCtx.UserModel.FindByPhone(l.ctx, in.Phone)
		if err == nil {
			userEntities = append(userEntities,userEntity)
		}
	}else if in.Name != "" {
		userEntities,err = l.svcCtx.UserModel.ListByName(l.ctx,in.Name)
	}else if len(in.Ids) > 0 {
		userEntities,err = l.svcCtx.UserModel.ListByIds(l.ctx,in.Ids)
	}
	if err != nil && err != sqlc.ErrNotFound {
		//fmt.Printf("FindUser err: %v \n",err)
		return nil, err
	}
	// 确保即使没有找到记录也返回空切片而非nil
	if userEntities == nil {
		userEntities = []*models.Users{}
	}
	
	var resp []*user.UserEntity
	copier.Copy(&resp, &userEntities)
	fmt.Printf("userEntities: %v \n",userEntities)
	fmt.Printf("resp: %v \n",resp)
	return &user.FindUserResp{
		User: resp,
	}, nil
}
