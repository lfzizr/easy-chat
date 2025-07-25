package logic

import (
	"context"

	"gitee.com/lfzizr/easy-chat/apps/social/rpc/internal/svc"
	"gitee.com/lfzizr/easy-chat/apps/social/rpc/social"
	"gitee.com/lfzizr/easy-chat/apps/social/socialmodels"
	"gitee.com/lfzizr/easy-chat/pkg/constants"
	"gitee.com/lfzizr/easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)
var (
	ErrFriendReqPassed = xerr.NewMsg("好友申请已通过")
	ErrFriendReqRefused = xerr.NewMsg("好友申请已拒绝")
)
type FriendPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInHandleLogic) FriendPutInHandle(in *social.FriendPutInHandleReq) (*social.FriendPutInHandleResp, error) {
	// todo: add your logic here and delete this line
	//获取好友申请记录
	friendReq, err := l.svcCtx.FriendRequestsModel.FindOne(l.ctx, uint64(in.FriendReqId))
	if err != nil {
		return nil,errors.Wrapf(xerr.NewDBErr(),"find friendReq err %v req %v",err,friendReq)
	}
	//验证是否有处理
	switch constants.HandleResult(friendReq.HandleResult.Int64) {
	case constants.PassHandlerResult:
		return nil,errors.WithStack(ErrFriendReqPassed)
	case constants.RefuseHandlerResult:
		return nil,errors.WithStack(ErrFriendReqRefused)
	}
	//修改handleResult 》通过建立两条好友记录关系 》 事务
	friendReq.HandleResult.Int64 = int64(in.HandleResult)
	err = l.svcCtx.Trans(l.ctx,func(ctx context.Context,session sqlx.Session) error {
		if err = l.svcCtx.FriendRequestsModel.Update(ctx,session,friendReq);err!= nil {
			return errors.Wrapf(xerr.NewDBErr(),"updata friendRequests err %v req %v",err,friendReq)
		}
		if constants.HandleResult(in.HandleResult) != constants.PassHandlerResult {
			return nil
		}
		friends :=[]*socialmodels.Friends{
			{
				UserId: friendReq.UserId,
				FriendUid: friendReq.ReqUid,
			},
			{
				UserId: friendReq.ReqUid,
				FriendUid: friendReq.UserId,
			},
		}
		if _,err := l.svcCtx.FriendsModel.Inserts(ctx,session,friends...);err != nil {
			return errors.Wrapf(xerr.NewDBErr(),"inserts friends err %v req %v",err,friendReq)
		}
		return nil
	})
	return &social.FriendPutInHandleResp{}, nil
}
