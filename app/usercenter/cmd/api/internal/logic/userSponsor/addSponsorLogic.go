package userSponsor

import (
	"github.com/pkg/errors"
	"context"
	"github.com/jinzhu/copier"
	"lottery/app/usercenter/cmd/rpc/pb"
	"lottery/common/ctxdata"
	"lottery/common/xerr"
	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSponsorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加 抽奖发起人（赞助商）
func NewAddSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSponsorLogic {
	return &AddSponsorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddSponsorLogic.AddSponsor 添加赞助商
// 该方法接收一个创建赞助商的请求，并返回创建结果或错误
// 主要步骤包括：将请求数据复制到protobuf结构体、获取用户ID、调用RPC服务添加赞助商信息
func (l *AddSponsorLogic) AddSponsor(req *types.CreateSponsorReq) (resp *types.CreateSponsorResp, err error) {
    // 创建一个protobuf结构体实例，用于传递给RPC服务
    pbSponsor := new(pb.AddUserSponsorReq)

    // 将请求数据复制到protobuf结构体中，以便于在RPC调用中使用
    err = copier.Copy(pbSponsor, req)
    if err != nil {
        // 如果数据复制失败，则返回错误
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to copy data:%+v,err:%s", req, err)
    }

    // 从上下文中获取当前用户的ID，并将其设置为赞助商的UserID
    pbSponsor.UserId = ctxdata.GetUidFromCtx(l.ctx)

    // 调用RPC服务，添加用户赞助商信息
    sponsor, err := l.svcCtx.UsercenterRpc.AddUserSponsor(l.ctx, pbSponsor)
    if err != nil {
        // 如果添加赞助商失败，则返回错误
        return nil, errors.Wrapf(xerr.NewErrMsg("failed to add sponsor"), "failed to add sponsor:%+v,err:%s", sponsor, err)
    }

    // 返回创建赞助商的响应，包括新创建的赞助商ID
    return &types.CreateSponsorResp{
        Id: sponsor.Id,
    }, nil
}

