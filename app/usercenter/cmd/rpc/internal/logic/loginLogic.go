package logic

import (
	"context"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"lottery/app/usercenter/model"
	"lottery/common/tool"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/rpc/internal/svc"
	"lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自定义的服务
// 用户登录
func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var userId int64
	var err error
	// 根据认证类型选择不同的登录方式
	switch in.AuthType {
	case model.UserAuthTypeSystem:
		// 使用手机号和密码登录
		userId, err = l.loginByMobile(in.AuthKey, in.Password)
	default:
		// 如果认证类型不支持，则返回通用错误
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	if err != nil {
		// 如果登录失败，则返回错误
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
		return nil, errors.Wrapf(ErrGenerateTokenError, "生成 Token 失败, 错误: %v, user_id: %d", err, userId)
	}
	// 返回登录响应
	return &pb.LoginResp{
		AccessToken:   tokenResp.AccessToken,
		AccessExpire:  tokenResp.AccessExpire,
		RefreshAfter:  tokenResp.RefreshAfter,
	}, nil
}

// 使用手机号和密码登录
func (l *LoginLogic) loginByMobile(mobile string, password string) (int64, error) {
	// 根据手机号查找用户
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && err != model.ErrNotFound {
		// 如果查找失败且错误不是 ErrNotFound，则包装错误并返回
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据手机号查询用户信息失败, 手机号: %s, 错误: %v", mobile, err)
	}
	if user == nil {
		// 如果用户不存在，则包装错误并返回
		return 0, errors.Wrapf(ErrUserNoExistsError, "用户不存在, 手机号: %s", mobile)
	}
	// 验证密码是否正确
	if !(tool.Md5ByString(password) == user.Password) {
		return 0, errors.Wrap(ErrUsernamePwdError, "密码错误")
	}
	// 返回用户ID
	return user.Id, nil
}

// 使用小程序微信登录
func (l *LoginLogic) loginBySmallWx() error {
	return nil
}

