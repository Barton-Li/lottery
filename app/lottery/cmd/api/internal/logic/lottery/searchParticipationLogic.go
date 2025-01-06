package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"lottery/app/lottery/cmd/rpc/lottery"
	"lottery/app/usercenter/cmd/rpc/usercenter"

	"lottery/app/lottery/cmd/api/internal/svc"
	"lottery/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchParticipationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchParticipationLogic {
	return &SearchParticipationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SearchParticipation 根据要求搜索抽奖抽奖参与信息。
// req: 搜索抽奖参与的请求参数，包括抽奖ID、页面大小和页面索引。
// 返回: 搜索结果和错误信息，搜索结果包括参与用户信息和总数。
func (l *SearchParticipationLogic) SearchParticipation(req *types.SearchLotteryParticipationReq) (resp *types.SearchLotteryParticipationResp, err error) {
	// 调用RPC服务搜索抽奖参与信息。
	res, err := l.svcCtx.LotteryRpc.SearchLotteryParticipation(l.ctx, &lottery.SearchLotteryParticipationReq{
		LotteryId: req.LotteryId,
		PageSize:  req.PageSize,
		PageIndex: req.PageIndex,
	})
	if err != nil {
		return nil, err
	}

	// 提取参与用户的ID。
	var userIds []int64
	for i := range res.List {
		userIds = append(userIds, res.List[i].UserId)
	}

	// 根据用户ID获取用户信息。
	userInfos := new(usercenter.GetUserInfoByUserIdsResp)
	if len(userIds) > 0 {
		userInfos, err = l.svcCtx.UsercenterRpc.GetUserInfoByUserIds(l.ctx, &usercenter.GetUserInfoByUserIdsReq{
			UserIds: userIds,
		})
		if err != nil {
			return nil, err
		}
	}

	// 对用户昵称进行脱敏处理。
	for idx, item := range userInfos.UserInfo {
		if len(item.Nickname) > 2 {
			item.Nickname = item.Nickname[:1] + "**" + item.Nickname[len(item.Nickname)-1:]
		} else {
			item.Nickname = item.Nickname[:] + "**"
		}
		userInfos.UserInfo[idx] = item
	}

	// 初始化响应对象并复制用户信息到响应列表。
	resp = new(types.SearchLotteryParticipationResp)
	err = copier.Copy(&resp.List, userInfos.UserInfo)
	if err != nil {
		return nil, err
	}

	// 设置参与用户总数。
	resp.Count = res.Count

	// 返回搜索结果和错误信息。
	return
}
