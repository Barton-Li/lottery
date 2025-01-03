package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2/cache"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"lottery/common/tool"
	"lottery/common/xerr"
	usercenterModel "lottery/app/usercenter/model"
	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/zeromicro/go-zero/core/logx"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"strings"
)

// ErrWxMiniAuthFailError error
var ErrWxMiniAuthFailError = xerr.NewErrMsg("wechat mini auth fail")

type WxMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// WxMiniAuthLogic 是一个处理微信小程序认证逻辑的结构体。
// 该方法负责处理微信小程序的认证请求，包括验证用户身份和解密用户数据。
func (l *WxMiniAuthLogic) WxMiniAuth(req *types.WXMinAuthReq) (resp *types.WXMinAuthResp, err error) {
    // 初始化微信小程序客户端，使用配置中的AppID和AppSecret。
    minProgram := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
        AppID:     l.svcCtx.Config.WxMiniConf.AppId,
        AppSecret: l.svcCtx.Config.WxMiniConf.AppSecret,
        Cache:     cache.NewMemory(),
    })

    // 使用微信小程序的Code2Session接口进行授权。
    authResult, err := minProgram.GetAuth().Code2Session(req.Code)
    if err != nil || authResult.ErrCode != 0 || authResult.OpenID == "" {
        // 如果授权失败，记录错误日志并返回错误信息。
        logx.Error("wechat mini auth fail, err: %v", err)
        return nil, errors.Wrapf(xerr.NewErrMsg("wechat mini auth fail"), "发起授权请求失败 err: %v,code: %s,authResult: %v", err, req.Code, authResult)
    }

    // 解密用户数据，包括昵称和头像URL。
    userData, err := minProgram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
    if err != nil {
        // 如果解密失败，返回错误信息。
        return nil, errors.Wrapf(xerr.NewErrMsg("解析数据失败"), "解析数据失败 err: %v,userData: %v,authResult: %v", err, userData, authResult)
    }

    // 尝试从用户中心服务获取用户认证信息。
    var userId int64
    rpcResp, err := l.svcCtx.UsercenterRpc.GetUserAuthByAuthKey(l.ctx, &usercenter.GetUserAuthByAuthKeyReq{
        AuthType: usercenterModel.UserAuthTypeMPWX,
        AuthKey:  authResult.OpenID,
    })
    if err != nil {
        // 如果RPC调用失败，返回错误信息。
        return nil, errors.Wrapf(xerr.NewErrMsg("rpc call userAuthByAuthKey fail"), "rpc call userAuthByAuthKey fail, err: %v,arrResult: %+v", err, authResult)
    }

    // 如果用户不存在，则生成新的AccessToken。
    if rpcResp.UserAuth == nil || rpcResp.UserAuth.Id == 0 {
        // 如果是新用户，生成默认昵称和头像，并进行注册。
        if len(req.Nickname) == 0 {
            nicknameArr := []string{userData.NickName, tool.Krand(6, tool.KC_RAND_KIND_NUM)}
            nickname := strings.Join(nicknameArr, "")
            req.Nickname = nickname
        }
        if len(req.Avatar) == 0 {
            req.Avatar = userData.AvatarURL
        }
        openId := authResult.OpenID
        wxMiniRegisterResp, err := l.svcCtx.UsercenterRpc.WxMiniRegister(l.ctx, &usercenter.WXMiniRegisterReq{
            AuthKey:  openId,
            AuthType: usercenterModel.UserAuthTypeMPWX,
            Nickname: req.Nickname,
            Avatar:   req.Avatar,
        })
        if err != nil {
            // 如果注册失败，返回错误信息。
            return nil, errors.Wrapf(xerr.NewErrMsg("UsercenterRpc.Register err"), "UsercenterRpc.Register err :%v, authResult : %+v", err, authResult)
        }
        // 返回注册成功后的AccessToken信息。
        return &types.WXMinAuthResp{
            AccessToken:       wxMiniRegisterResp.AccessToken,
            AccessTokenExpire: wxMiniRegisterResp.AccessExpire,
            RefreshAfter:      wxMiniRegisterResp.RefreshAfter,
        }, nil
    } else {
        // 如果用户已存在，直接生成新的AccessToken。
        userId = rpcResp.UserAuth.UserId
        tokenResp, err := l.svcCtx.UsercenterRpc.GenerateToken(l.ctx, &usercenter.GenerateTokenReq{
            UserId: userId,
        })
        if err != nil {
            // 如果生成Token失败，返回错误信息。
            return nil, errors.Wrapf(xerr.NewErrMsg("usercenterRpc.GenerateToken err"), "usercenterRpc.GenerateToken err :%v, userId : %d", err, userId)

        }
        // 返回生成的AccessToken信息。
        return &types.WXMinAuthResp{
            AccessToken:       tokenResp.AccessToken,
            AccessTokenExpire: tokenResp.AccessExpire,
            RefreshAfter:      tokenResp.RefreshAfter,
        }, nil
    }
}

