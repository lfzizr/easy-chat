package logic

import (
	"context"
	"database/sql"
	"time"

	"gitee.com/lfzizr/easy-chat/apps/social/rpc/internal/svc"
	"gitee.com/lfzizr/easy-chat/apps/social/rpc/social"
	"gitee.com/lfzizr/easy-chat/apps/social/socialmodels"
	"gitee.com/lfzizr/easy-chat/pkg/constants"
	"gitee.com/lfzizr/easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	// todo: add your logic here and delete this line
	//查看和目标用户是否已经是好友
	friends, err := l.svcCtx.FriendsModel.FindByUidAndFid(l.ctx, in.UserId, in.ReqUid)
	if err != nil && err != socialmodels.ErrNotFound{
		return nil,errors.Wrapf(xerr.NewDBErr(),"find friends by uid and fid err %v req %v",err,in)
	}
	if friends != nil {
		return &social.FriendPutInResp{},errors.Wrapf(xerr.NewDBErr(),"friend already exist req %v",in)
	}

	//是否有过好友申请
	friendsReqs,err := l.svcCtx.FriendRequestsModel.FindByReqUidAndUserId(l.ctx, in.ReqUid, in.UserId)
	if err != nil && err != socialmodels.ErrNotFound{
		return nil,errors.Wrapf(xerr.NewDBErr(),"find friendsRequests by req uid and user id err %v req %v",err,in)
	}
	if friendsReqs != nil {
		return &social.FriendPutInResp{},err
	}
	//插入申请记录
	_,err = l.svcCtx.FriendRequestsModel.Insert(l.ctx,&socialmodels.FriendRequests{
		ReqUid: in.ReqUid,
		UserId: in.UserId,
		ReqMsg: sql.NullString{
			Valid: true,
			String: in.ReqMsg,
		},
		ReqTime: time.Unix(in.ReqTime,0),
		HandleResult: sql.NullInt64{
			Valid: true,
			Int64: int64(constants.NoHandlerResult),
		},
	})
	if err != nil {
		return nil,errors.Wrapf(xerr.NewDBErr(),"insert friendRequests err %v req %v",err,in)
	}
	return &social.FriendPutInResp{}, nil
}
