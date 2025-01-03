package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/model"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"
	"lottery/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserAddressLogic {
	return &SearchUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchUserAddress 根据用户请求搜索用户地址信息。
// in 参数为搜索用户地址的请求，包含分页信息。
// 返回用户地址列表和潜在的错误。
func (l *SearchUserAddressLogic) SearchUserAddress(in *pb2.SearchUserAddressReq) (*pb2.SearchUserAddressResp, error) {
	// 调用模型层方法获取用户地址列表。
	list, err := l.svcCtx.UserAddressModel.List(l.ctx, in.Page, in.Limit)
	// 检查是否有错误发生，如果发生错误且不是未找到资源的错误，则返回错误。
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get user's homestay order err : %v , in :%+v", err, in)
	}
	// 初始化响应列表。
	var resp []*pb2.UserAddress
	// 如果获取到的地址列表长度大于0，则遍历列表，将每个地址信息转换为pb2.UserAddress类型，并添加到响应列表中。
	if len(list) > 0 {
		for _, address := range list {
			var pbAddress pb2.UserAddress
			// 使用copier库复制地址信息，避免手动赋值的繁琐。
			_ = copier.Copy(&pbAddress, address)
			resp = append(resp, &pbAddress)
		}
	}
	// 返回用户地址响应对象，包含用户地址列表。
	return &pb2.SearchUserAddressResp{
		UserAddress: resp,
	}, nil
}
