package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/common/xerr"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLotteryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryDetailLogic {
	return &LotteryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LotteryDetail 获取抽奖详情信息。
// 该方法接收一个 LotteryDetailReq 请求对象作为输入，其中包含抽奖ID和用户ID，
// 并返回一个包含抽奖详情和用户参与状态的 LotteryDetailResp 对象，或错误信息。
func (l *LotteryDetailLogic) LotteryDetail(in *pb.LotteryDetailReq) (*pb.LotteryDetailResp, error) {
    // 提取请求中的抽奖ID。
    lotteryID:=in.Id
    
    // 通过抽奖ID查找对应的奖品信息。
    prize, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, lotteryID)
    if err != nil {
        // 如果查找奖品信息出错，返回错误。
        return nil, err
    }
    
    // 查找抽奖详情信息。
    lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, lotteryID)
    if err != nil {
        // 如果查找抽奖信息出错，包装错误信息并返回。
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_LOTTERY_BYLOTTERYID_ERROR), "query lottery by lotteryId:%d error:%s,resp:%v", lotteryID, err, lottery)
    }
    
    // 初始化抽奖详情响应对象。
    resp:=new(pb.LotteryDetailResp)
    resp.Lottery=new(pb.Lottery)
    
    // 将数据库中的抽奖信息复制到响应对象中。
    _ = copier.Copy(resp.Lottery, lottery)
    
    // 将时间字段转换为Unix时间戳格式。
    resp.Lottery.AnnounceTime = lottery.AnnounceTime.Unix()
    resp.Lottery.PubilshTime = lottery.PublishTime.Time.Unix()
    resp.Lottery.AwardDeadline = lottery.AwardDeadline.Unix()
    resp.Lottery.CreateTime = lottery.CreateTime.Unix()
    resp.Lottery.UpdateTime = lottery.UpdateTime.Unix()
    
    // 将奖品信息复制到响应对象中。
    copier.Copy(&resp.Prizes, &prize)
    
    // 检查用户是否参与了当前抽奖活动。
    count,err:=l.svcCtx.LotteryParticipationModel.CheckIsParticipatedByUserIdAndLotteryId(l.ctx,in.UserId,lotteryID)
    if err!=nil {
        // 如果检查用户参与状态出错，返回错误。
        return nil, err
    }
    
    // 记录日志，显示用户是否参与了当前抽奖。
    l.Logger.Debugf("获取当前用户是否参与当前抽奖：%v",count)
    
    // 设置响应对象中的用户参与状态字段。
    resp.IsParticipated=count
    
    // 返回抽奖详情响应对象。
    return resp, nil
}

