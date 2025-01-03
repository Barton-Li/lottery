package userContact

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/pb"
	"lottery/common/ctxdata"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加抽奖发起人的联系方式
func NewAddContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddContactLogic {
	return &AddContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddContactLogic.AddContact 添加联系人
// 该方法接收一个创建联系人的请求，并返回创建结果或错误
// 主要步骤包括：将请求转换为protobuf格式，从上下文中获取用户ID，
// 序列化联系人内容，并调用远程服务完成联系人添加操作
func (l *AddContactLogic) AddContact(req *types.CreateContactReq) (resp *types.CreateContactResp, err error) {
    // 创建一个protobuf格式的添加联系人请求对象
    pbContactReq := new(pb.AddUserContactReq)

    // 将传入的请求参数复制到protobuf请求对象中
    err = copier.Copy(pbContactReq, req)
    if err != nil {
        // 如果复制失败，则返回错误信息，包装原始错误信息以提供更多上下文
        return nil, errors.Wrapf(xerr.NewErrMsg("failed to add contact"), "failed to add contact:%+v,err:%s", req, err)
    }

    // 从上下文中获取当前用户的ID，并设置到请求对象中
    pbContactReq.UserId = ctxdata.GetUidFromCtx(l.ctx)

    // 将联系人内容序列化为JSON格式
    ContentByte, err := json.Marshal(req.Content)
    if err != nil {
        // 如果序列化失败，则返回错误信息，包装原始错误信息以提供更多上下文
        return nil, errors.Wrapf(xerr.NewErrMsg("failed to add contact"), "failed to add contact:%+v, err:%s", req, err)
    }

    // 将序列化的联系人内容设置到请求对象中
    pbContactReq.Content = string(ContentByte)

    // 调用远程服务，尝试添加联系人
    contact, err := l.svcCtx.UsercenterRpc.AddUserContact(l.ctx, pbContactReq)
    if err != nil {
        // 如果添加失败，则返回错误信息，包装原始错误信息以提供更多上下文
        return nil, errors.Wrapf(xerr.NewErrMsg("failed to add contact"), "failed to add contact:%+v, err:%s", req, err)
    }

    // 如果添加成功，则构建响应对象并返回
    return &types.CreateContactResp{
        Id: contact.Id,
    }, nil
}

