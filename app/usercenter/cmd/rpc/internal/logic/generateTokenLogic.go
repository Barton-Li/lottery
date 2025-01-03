package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"lottery/common/ctxdata"
	"lottery/common/xerr"
	"time"

	"lottery/app/usercenter/cmd/rpc/internal/svc"
	"lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
    // 获取当前时间的时间戳（单位为秒）
    now := time.Now().Unix()

    // 从配置中读取访问令牌的过期时间（单位为秒）
    accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire

    // 调用getJwtToken方法生成访问令牌
    // 参数包括访问令牌的秘密密钥、当前时间戳、访问令牌的过期时间和用户ID
    accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId)

    // 如果生成访问令牌时出错，返回错误信息
    if err != nil {
        return nil, errors.Wrapf(ErrGenerateTokenError, "生成访问令牌错误:%v, 用户ID:%d", err, in.UserId)
    }

    // 如果生成成功，返回包含访问令牌、访问令牌过期时间和刷新时间的响应
    return &pb.GenerateTokenResp{
        AccessToken:  accessToken, // 访问令牌字符串
        AccessExpire: now + accessExpire, // 访问令牌的过期时间戳
        RefreshAfter: now + accessExpire/2, // 建议的刷新访问令牌的时间戳（过期时间的一半）
    }, nil
}


// 生成 JWT Token
func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	// 创建一个 jwt.MapClaims 对象，用于存储 JWT 的声明信息
	claims := make(jwt.MapClaims)
	// 设置 Token 的过期时间 (exp)，即当前时间 (iat) 加上 seconds
	claims["exp"] = iat + seconds
	// 设置 Token 的签发时间 (iat)
	claims["iat"] = iat
	// 设置 Token 中的用户ID
	claims[ctxdata.CtxKeyJwtUserId] = userId
	// 创建一个新的 JWT Token，使用 HS256 签名方法
	token := jwt.New(jwt.SigningMethodHS256)
	// 将声明信息设置到 Token 中
	token.Claims = claims
	// 使用 secretKey 对 Token 进行签名，并返回签名后的 Token 字符串
	return token.SignedString([]byte(secretKey))
}
