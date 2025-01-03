package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"lottery/app/usercenter/model"
	"lottery/common/tool"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/rpc/internal/svc"
	"lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserAlreadyRegisterError = xerr.NewErrMsg("用户已经注册")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册用户
func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 根据手机号查找用户
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && err != model.ErrNotFound {
		// 如果查找失败且错误不是 ErrNotFound，则包装错误并返回
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "手机号:%s 错误:%v", in.Mobile, err)
	}
	if user != nil {
		// 如果用户已存在，则包装错误并返回
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "用户已注册 手机号:%s, 错误:%v", in.Mobile, err)
	}
	var userId int64
	// 开始事务处理
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := new(model.UserInfo)
		user.Mobile = in.Mobile
		// 如果昵称为空，则生成一个随机昵称
		if len(user.Nickname) == 0 {
			user.Nickname = tool.Krand(8, tool.KC_RAND_KIND_ALL)
		}
		// 如果密码不为空，则对密码进行 MD5 加密
		if len(in.Password) > 0 {
			user.Password = tool.Md5ByString(in.Password)
		}
		// 插入用户信息到数据库
		insertResult, err := l.svcCtx.UserModel.Insert(ctx, user)
		if err != nil {
			// 如果插入失败，则包装错误并返回
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "注册插入用户信息错误: %v, 用户信息: %+v", err, user)
		}
		// 获取插入的用户ID
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			// 如果获取用户ID失败，则包装错误并返回
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "注册获取最后插入ID错误: %v, 用户信息: %+v", err, user)
		}
		userId = lastId
		userAuth := new(model.UserAuth)
		userAuth.UserId = lastId
		userAuth.AuthKey = in.AuthKey
		userAuth.AuthType = in.AuthType
		// 插入用户认证信息到数据库
		if _, err := l.svcCtx.UserAuthModel.Insert(ctx, session, userAuth); err != nil {
			// 如果插入失败，则包装错误并返回
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "注册插入用户认证信息错误: %v, 用户认证信息: %+v", err, userAuth)
		}
		userSponsor := new(model.UserSponsor)
		userSponsor.UserId = lastId
		userSponsor.IsShow = 1
		userSponsor.Name = tool.Krand(8, tool.KC_RAND_KIND_ALL)
		userSponsor.Avatar = "https://img01.yzcdn.cn/vant/cat.jpeg"
		// 插入用户赞助信息到数据库
		if _, err := l.svcCtx.UserSponsorModel.Insert(ctx, userSponsor); err != nil {
			// 如果插入失败，则包装错误并返回
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "注册插入用户赞助信息错误: %v, 用户赞助信息: %+v", err, userSponsor)
		}
		return nil

	}); err != nil {
		// 如果事务处理失败，记录错误日志
		logx.Error("注册错误: %v", err)
		return nil, err
	}
	// 创建生成 Token 的逻辑实例
	genTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	// 生成 Token
	tokenResp, err := genTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		// 如果生成 Token 失败，则包装错误并返回
		return nil, errors.Wrapf(ErrGenerateTokenError, "生成 Token 错误 user_id: %d, 错误: %v", userId, err)
	}

	// 返回注册响应
	return &pb.RegisterResp{
		AccessExpire: tokenResp.AccessExpire,
		AccessToken:  tokenResp.AccessToken,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

