package userSponsor

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/common/ctxdata"
	"lottery/common/xerr"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 我的赞助商列表（赞助商）
func NewSponsorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorListLogic {
	return &SponsorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SponsorList 获取赞助者列表
// 该方法通过调用远程服务SearchUserSponsor来获取用户的赞助者列表，并将结果转换为指定的响应类型返回。
// 参数 req: 包含分页信息和用户ID的请求对象，用于指定要获取的赞助者列表的页码和每页的条数。
// 返回值 resp: 包含赞助者列表的响应对象。
// 返回值 err: 如果在处理过程中遇到错误，则返回错误信息。
func (l *SponsorListLogic) SponsorList(req *types.SponsorListReq) (resp *types.SponsorListResp, err error) {
    // 调用远程服务SearchUserSponsor，传递分页信息和用户ID，获取赞助者列表。
    rpcSponsorList, err := l.svcCtx.UsercenterRpc.SearchUserSponsor(l.ctx, &usercenter.SearchUserSponsorReq{
        Page:   req.Page,
        Limit:  req.PageSize,
        UserId: ctxdata.GetUidFromCtx(l.ctx),
    })
    // 如果调用远程服务失败，则返回错误信息。
    if err != nil {
        return nil, errors.Wrapf(xerr.NewErrMsg("get sponsor list fail"), "Failed to get Sponsor:%+v,err:%s", rpcSponsorList, err)
    }
    // 初始化赞助者列表。
    var SponsorList []types.Sponsor
    // 如果远程服务返回的赞助者列表不为空，则遍历列表，将每个赞助者的信息复制到新的类型中，并添加到赞助者列表中。
    if len(rpcSponsorList.UserSponsor) > 0 {
        for _, item := range rpcSponsorList.UserSponsor {
            var t types.Sponsor
            // 使用copier库将item的值复制给t，实现类型转换。
            _ = copier.Copy(&t, item)
            // 将转换后的赞助者信息添加到列表中。
            SponsorList = append(SponsorList, t)
        }
    }
    // 返回包含赞助者列表的响应对象。
    return &types.SponsorListResp{
        List: SponsorList,
    }, nil
}

