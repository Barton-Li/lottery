package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"lottery/app/usercenter/model"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/rpc/internal/svc"
	"lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxMiniRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWxMiniRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniRegisterLogic {
	return &WxMiniRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// WxMiniRegister 实现小程序用户注册逻辑。
// 该方法首先在数据库中创建用户信息，然后创建用户认证信息和用户赞助商信息。
// 最后，生成用户令牌并返回。
// 参数:
//   in *pb.WXMiniRegisterReq - 包含小程序用户注册所需信息的请求对象。
// 返回值:
//   *pb.WXMiniRegisterResp - 包含用户令牌信息的响应对象。
//   error - 错误信息，如果注册过程中发生错误。
func (l *WxMiniRegisterLogic) WxMiniRegister(in *pb.WXMiniRegisterReq) (*pb.WXMiniRegisterResp, error) {
    var userId int64

    // 使用事务处理用户注册过程，确保数据一致性。
    if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
        user := new(model.UserInfo)
        user.Nickname = in.Nickname
        user.Avatar = in.Avatar

        // 插入用户信息到数据库。
        transInsert, err := l.svcCtx.UserModel.TransInsert(ctx, session, user)
        if err != nil {
            return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "WxMiniRegister insert error:%v,user%+v", err, user)
        }

        // 获取插入用户的ID。
        lastId, err := transInsert.LastInsertId()
        if err != nil {
            return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "WxMiniRegister LastInsertId error:%v,user:%+v", err, user)
        }
        userId = lastId

        // 创建并插入用户认证信息。
        userAuth := new(model.UserAuth)
        userAuth.UserId = lastId
        userAuth.AuthKey = in.AuthKey
        userAuth.AuthType = in.AuthType
        if _, err := l.svcCtx.UserAuthModel.Insert(ctx, session, userAuth); err != nil {
            return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "WxMiniRegister insert error:%v, userAuth:%+v", err, userAuth)
        }

        // 创建并插入用户赞助商信息。
        userSponsor := new(model.UserSponsor)
        userSponsor.UserId = lastId
        userSponsor.Avatar = in.Avatar
        userSponsor.IsShow = 1
        userSponsor.Name = in.Nickname
        if _, err := l.svcCtx.UserSponsorModel.Insert(ctx, userSponsor); err != nil {
            return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "WxMiniRegister insert error:%v, userSponsor:%+v", err, userSponsor)
        }

        return nil
    }); err != nil {
        logx.Error("trans error:%v", err)
        return nil, err
    }

    // 生成用户令牌。
    generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
    tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
        UserId: userId,
    })
    if err != nil {
        return nil, errors.Wrapf(ErrGenerateTokenError, "WxMiniRegister GenerateToken error:%v,%d", err, userId)
    }

    // 返回用户令牌信息。
    return &pb.WXMiniRegisterResp{
        AccessToken:  tokenResp.AccessToken,
        AccessExpire: tokenResp.AccessExpire,
        RefreshAfter: tokenResp.RefreshAfter,
    }, nil
}

