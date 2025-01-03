package address

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 收货地址列表
func NewAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressListLogic {
	return &AddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddressList 获取用户地址列表
// 该方法通过调用RPC服务获取用户的地址信息，并将其格式化后返回
// 参数:
//   req *types.AddressListReq: 包含分页信息的请求对象
// 返回值:
//   *types.AddressListResp: 包含用户地址列表的响应对象
//   error: 错误信息，如果执行过程中遇到错误则返回
func (l *AddressListLogic) AddressList(req *types.AddressListReq) (resp *types.AddressListResp, err error) {
    // 调用RPC服务，搜索用户地址
    rpcAddressList, err := l.svcCtx.UsercenterRpc.SearchUserAddress(l.ctx, &usercenter.SearchUserAddressReq{
        Page:  req.Page,
        Limit: req.PageSize,
    })
    // 如果发生错误，包装错误信息并返回
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get SearchLottery"), "req:%+v,err:%v", req, err)
    }

    // 初始化地址列表变量
    var AddressList []types.UserAddress
    // 如果RPC调用返回的地址列表不为空，则遍历每个地址并进行处理
    if len(rpcAddressList.UserAddress) > 0 {
        for _, item := range rpcAddressList.UserAddress {
            // 初始化用户地址对象
            var t types.UserAddress
            // 将RPC调用返回的地址信息复制到用户地址对象中
            _ = copier.Copy(&t, item)
            // 将地址中的地区信息从JSON字符串解析为对象
            _ = json.Unmarshal([]byte(item.District), &t.District)
            // 将处理后的地址对象添加到地址列表中
            AddressList = append(AddressList, t)
        }
    }

    // 返回包含地址列表的响应对象
    return &types.AddressListResp{
        List: AddressList,
    }, nil
}
