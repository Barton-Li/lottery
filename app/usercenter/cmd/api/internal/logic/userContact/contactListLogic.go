package userContact

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"lottery/common/ctxdata"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContactListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 抽奖发起人的联系方式列表
func NewContactListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContactListLogic {
	return &ContactListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ContactList 获取用户联系人列表
// 该方法通过调用远程服务搜索用户联系人信息，并将其解析为联系人列表响应对象
func (l *ContactListLogic) ContactList(req *types.ContactListReq) (resp *types.ContactListResp, err error) {
    // 调用远程服务获取用户联系人信息
    rpcContactList, err := l.svcCtx.UsercenterRpc.SearchUserContact(l.ctx, &usercenter.SearchUserContactReq{
        Page:   req.Page,
        Limit:  req.PageSize,
        UserId: ctxdata.GetUidFromCtx(l.ctx),
    })
    // 错误处理：如果远程服务调用失败，则返回错误
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrMsg("get contact list fail"), "Failed to get Contact:%+v,err:%s", rpcContactList, err)
    }

    // 初始化联系人列表变量
    var ContactList []types.Contact
    // 如果从远程服务获取到用户联系人信息，则遍历并解析每个联系人信息
    if len(rpcContactList.UserContact) > 0 {
        for _, contact := range rpcContactList.UserContact {
            var t types.Contact
            // 将远程服务返回的联系人信息复制到本地联系人对象
            _ = copier.Copy(&t, contact)
            // 解析联系人内容（假设内容是JSON格式）
            _ = json.Unmarshal([]byte(contact.Content), &t.Content)
            // 将解析后的联系人添加到联系人列表
            ContactList = append(ContactList, t)
        }
    }

    // 构建联系人列表响应对象
    resp = &types.ContactListResp{
        List: ContactList,
    }

    // 返回联系人列表响应对象和nil错误
    return resp, nil
}

