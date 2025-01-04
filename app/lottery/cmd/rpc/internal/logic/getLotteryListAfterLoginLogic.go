package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryListAfterLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryListAfterLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryListAfterLoginLogic {
	return &GetLotteryListAfterLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetLotteryListAfterLoginLogic 获取用户登录后的抽奖列表。
// 该方法根据用户的参与情况和请求参数，返回一个经过筛选和排序的抽奖列表。
// 参数:
//   in *pb.GetLotteryListAfterLoginReq - 包含用户ID、最后一条记录的ID、请求的列表大小等信息的请求对象。
// 返回值:
//   *pb.GetLotteryListAfterLoginResp - 包含筛选后的抽奖列表的响应对象。
//   error - 错误信息，如果执行过程中遇到任何问题，则返回相应的错误。
func (l *GetLotteryListAfterLoginLogic) GetLotteryListAfterLogin(in *pb.GetLotteryListAfterLoginReq) (*pb.GetLotteryListAfterLoginResp, error) {
	// 获取用户参与过的抽奖ID列表。
	ParticipatedLotteryIds, err := l.svcCtx.LotteryParticipationModel.GetParticipatedLotteryIdsByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	// 如果请求参数中的LastId为0，则获取最新的抽奖ID并设置为LastId。
	if in.LastId == 0 {
		lastId, err := l.svcCtx.LotteryModel.GetLastId(l.ctx)
		if err != nil {
			return nil, err
		}
		in.LastId = lastId + 1
	}

	// 确保ParticipatedLotteryIds不为空，以便后续使用。
	if ParticipatedLotteryIds == nil {
		ParticipatedLotteryIds = []int64{}
	}

	// 根据请求参数和用户参与情况，获取筛选后的抽奖列表。
	list, err := l.svcCtx.LotteryModel.GetLotteryListAfterLogin(l.ctx, in.Size, in.IsSelected, in.LastId, ParticipatedLotteryIds)
	if err != nil {
		return nil, err
	}

	// 初始化响应对象。
	var resp []*pb.Lottery
	if len(list) > 0 {
		for _, lottery := range list {
			// 将查询到的抽奖信息转换为pb.Lottery对象。
			var pbLottery pb.Lottery
			_ = copier.Copy(&pbLottery, lottery)
			pbLottery.PubilshTime = lottery.PublishTime.Time.Unix()
			pbLottery.AwardDeadline = lottery.AwardDeadline.Unix()
			pbLottery.Announcetype = lottery.AnnounceType
			pbLottery.AnnounceTime = lottery.AnnounceTime.Unix()
			pbLottery.IsAnnounced = lottery.IsAnnounced
			resp = append(resp, &pbLottery)
		}
	}

	// 返回包含筛选后抽奖列表的响应对象。
	return &pb.GetLotteryListAfterLoginResp{
		List: resp,
	}, nil
}
