package logic

import (
	"context"
	"lottery/app/usercenter/cmd/rpc/usercenter"
	"sort"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWonListByLotteryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}
type UserInfo struct {
	Id       int64
	Nickname string
	Avatar   string
}

func NewGetWonListByLotteryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWonListByLotteryIdLogic {
	return &GetWonListByLotteryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetWonListByLotteryId 根据抽奖活动ID获取中奖用户列表。
// 该方法首先查询参与并中奖的用户，然后根据中奖用户ID获取用户信息，
// 并将这些信息与奖品信息一起组织成响应返回。
func (l *GetWonListByLotteryIdLogic) GetWonListByLotteryId(in *pb.GetWonListByLotteryIdReq) (*pb.GetWonListByLotteryIdResp, error) {
	// 构建查询语句，筛选出指定抽奖活动且中奖的参与记录。
	builder := l.svcCtx.LotteryParticipationModel.SelectBuilder().Where("lottery_id=? and  is_won=1", in.LotteryId)
	participations, err := l.svcCtx.LotteryParticipationModel.FindAll(l.ctx, builder, "")
	if err != nil {
		return nil, err
	}

	// 提取中奖用户的ID，以便后续获取用户信息。
	userIds := make([]int64, 0)
	for _, participation := range participations {
		userIds = append(userIds, participation.UserId)
	}

	// 构建奖品与用户ID的映射，用于后续组织中奖信息。
	prizeUserMap := make(map[int64][]int64)
	for _, v := range participations {
		if _, ok := prizeUserMap[v.PrizeId]; ok {
			prizeUserMap[v.PrizeId] = make([]int64, 0)
		}
		prizeUserMap[v.PrizeId] = append(prizeUserMap[v.PrizeId], v.UserId)
	}

	// 如果没有中奖用户，直接返回空响应。
	if len(userIds) == 0 {
		return &pb.GetWonListByLotteryIdResp{}, nil
	}

	// 根据用户ID获取中奖用户的详细信息。
	userInfo, err := l.svcCtx.UserCenterRpc.GetUserInfoByUserIds(l.ctx, &usercenter.GetUserInfoByUserIdsReq{
		UserIds: userIds,
	})
	if err != nil {
		return nil, err
	}

	// 处理用户昵称，隐藏部分字符以保护用户隐私。
	for idx, item := range userInfo.UserInfo {
		if len(item.Nickname) > 2 {
			item.Nickname = item.Nickname[:1] + "**" + item.Nickname[len(item.Nickname)-1:]
		} else {
			item.Nickname = item.Nickname[:] + "**"
		}
		userInfo.UserInfo[idx] = item
	}

	// 将用户信息存储到映射中，便于后续查找。
	userInfoMap := make(map[int64]*UserInfo)
	for _, item := range userInfo.UserInfo {
		userInfoMap[item.Id] = &UserInfo{
			Avatar:   item.Avatar,
			Nickname: item.Nickname,
			Id:       item.Id,
		}
	}

	// 获取奖品信息，并按奖品级别排序。
	prizes, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, in.LotteryId)
	if err != nil {
		return nil, err
	}
	sort.Slice(prizes, func(i, j int) bool { return prizes[i].Level < prizes[j].Level })

	// 组织中奖信息列表，包括奖品信息和中奖用户信息。
	list := make([]*pb.WonList2, 0)
	for _, v := range prizes {
		prize := &pb.Prize{
			Id:        v.Id,
			LotteryId: v.LotteryId,
			Type:      v.Type,
			Name:      v.Name,
			Level:     v.Level,
			Thumb:     v.Thumb,
			Count:     v.Count,
			GrantType: v.GrantType,
		}
		users := make([]*pb.UserInfo, 0)
		for _, userId := range prizeUserMap[v.Id] {
			user := userInfoMap[userId]
			users = append(users, &pb.UserInfo{
				Id:       user.Id,
				Avatar:   []byte(user.Avatar),
				Nickname: []byte(user.Nickname),
			})
		}
		list = append(list, &pb.WonList2{
			Prize:       prize,
			Users:       users,
			WinnerCount: int64(len(prizeUserMap[v.Id])),
		})
	}

	// 返回组织好的中奖信息列表。
	return &pb.GetWonListByLotteryIdResp{
		List: list,
	}, nil
}
