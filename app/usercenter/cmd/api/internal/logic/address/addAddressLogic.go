package address

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

type AddAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加收货地址
func NewAddAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddressLogic {
	return &AddAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddAddress 用于添加用户地址。
// 该方法接收一个AddAddressReq请求对象，其中包含用户想要添加的地址信息。
// 它返回一个AddAddressResp响应对象，其中包含新添加地址的ID，以及可能发生的错误。
func (l *AddAddressLogic) AddAddress(req *types.AddAddressReq) (resp *types.AddAddressResp, err error) {
    // 创建一个protobuf请求对象，并将传入的请求数据复制到该对象中。
    pbAddressReq := new(pb.AddUserAddressReq)
    err = copier.Copy(pbAddressReq, req)
    if err != nil {
        return nil, err
    }

    // 从上下文中获取用户ID，并将其设置为地址请求的用户ID。
    pbAddressReq.UserId = ctxdata.GetUidFromCtx(l.ctx)

    // 将区县信息序列化为JSON格式，并将其设置为地址请求的区县信息。
    districtByte, err := json.Marshal(req.District)
    if err != nil {
        return nil, err
    }
    pbAddressReq.District = string(districtByte)

    // 调用RPC服务，添加用户地址。
    addAddress, err := l.svcCtx.UsercenterRpc.AddUserAddress(l.ctx, pbAddressReq)
    if err != nil {
        // 如果添加地址失败，返回一个错误信息。
        return nil, errors.Wrapf(xerr.NewErrMsg("add address fail"), "add address rpc AddUserAddress fail req %+v,err:%v", req, err)
    }

    // 返回添加地址的响应，其中包含新地址的ID。
    return &types.AddAddressResp{
        Id: addAddress.Id,
    }, nil
}

