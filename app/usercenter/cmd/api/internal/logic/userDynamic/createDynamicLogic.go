package userDynamic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/pb"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDynamicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建用户动态
func NewCreateDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDynamicLogic {
	return &CreateDynamicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateDynamic 用于创建用户动态。
// 该方法接收一个 CreateDynamicReq 请求对象，验证请求数据的完整性，
// 将请求数据转换为 protobuf 格式，然后调用远程服务添加用户动态。
// 如果请求数据不合法或远程服务调用失败，将返回相应的错误。
// 参数:
//   req (*types.CreateDynamicReq): 创建动态的请求对象，包含用户ID、动态URL和备注等信息。
// 返回值:
//   resp (*types.CreateDynamicResp): 创建动态的响应对象，包含新创建动态的ID。
//   err (error): 如果操作失败，将返回错误信息。
func (l *CreateDynamicLogic) CreateDynamic(req *types.CreateDynamicReq) (resp *types.CreateDynamicResp, err error) {
	// 初始化 protobuf 请求对象。
	pbDynamicReq := new(pb.AddUserDynamicReq)

	// 将传入的请求对象复制到 protobuf 请求对象中。
	err = copier.Copy(pbDynamicReq, req)
	if err != nil {
		return nil, errors.Wrapf(err, "copier.Copy error, req:%+v", req)
	}

	// 检查用户ID是否为空。
	if req.UserId == 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("user id is nil"), "user id is nil")
	}

	// 检查动态URL是否为空。
	if pbDynamicReq.DynamicUrl == "" {
		return nil, errors.Wrapf(xerr.NewErrMsg("dynamic url is nil"), "dynamic url is nil")
	}

	// 检查备注是否为空。
	if pbDynamicReq.Remark == "" {
		return nil, errors.Wrapf(xerr.NewErrMsg("remark is nil"), "remark is nil")
	}

	// 调用远程服务添加用户动态。
	addDynamic, err := l.svcCtx.UsercenterRpc.AddUserDynamic(l.ctx, pbDynamicReq)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to create dynamic"), "l.svcCtx.UsercenterRpc.AddUserDynamic error, req:%+v,err:%v", req, err)
	}

	// 返回创建动态的响应对象。
	return &types.CreateDynamicResp{
		Id: addDynamic.Id,
	}, nil
}
