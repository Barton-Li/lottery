// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: lottery.proto

package lottery

import (
	"context"

	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddClockTaskRecordReq                  = pb.AddClockTaskRecordReq
	AddClockTaskRecordResp                 = pb.AddClockTaskRecordResp
	AddLotteryParticipationReq             = pb.AddLotteryParticipationReq
	AddLotteryParticipationResp            = pb.AddLotteryParticipationResp
	AddLotteryReq                          = pb.AddLotteryReq
	AddLotteryResp                         = pb.AddLotteryResp
	AddPrizeReq                            = pb.AddPrizeReq
	AddPrizeResp                           = pb.AddPrizeResp
	AnnounceLotteryReq                     = pb.AnnounceLotteryReq
	AnnounceLotteryResp                    = pb.AnnounceLotteryResp
	CheckIsParticipatedReq                 = pb.CheckIsParticipatedReq
	CheckIsParticipatedResp                = pb.CheckIsParticipatedResp
	CheckSelectedLotteryParticipatedReq    = pb.CheckSelectedLotteryParticipatedReq
	CheckSelectedLotteryParticipatedResp   = pb.CheckSelectedLotteryParticipatedResp
	CheckUserCreateLotteryAndTodayReq      = pb.CheckUserCreateLotteryAndTodayReq
	CheckUserCreateLotteryAndTodayResp     = pb.CheckUserCreateLotteryAndTodayResp
	CheckUserCreateLotteryReq              = pb.CheckUserCreateLotteryReq
	CheckUserCreateLotteryResp             = pb.CheckUserCreateLotteryResp
	CheckUserCreatedLotteryAndThisWeekReq  = pb.CheckUserCreatedLotteryAndThisWeekReq
	CheckUserCreatedLotteryAndThisWeekResp = pb.CheckUserCreatedLotteryAndThisWeekResp
	CheckUserIsWonReq                      = pb.CheckUserIsWonReq
	CheckUserIsWonResp                     = pb.CheckUserIsWonResp
	ClockTask                              = pb.ClockTask
	CreatePrize                            = pb.CreatePrize
	DelLotteryReq                          = pb.DelLotteryReq
	DelLotteryResp                         = pb.DelLotteryResp
	DelPrizeReq                            = pb.DelPrizeReq
	DelPrizeResp                           = pb.DelPrizeResp
	GetLotteryByIdReq                      = pb.GetLotteryByIdReq
	GetLotteryByIdResp                     = pb.GetLotteryByIdResp
	GetLotteryListAfterLoginReq            = pb.GetLotteryListAfterLoginReq
	GetLotteryListAfterLoginResp           = pb.GetLotteryListAfterLoginResp
	GetLotteryListLastIdReq                = pb.GetLotteryListLastIdReq
	GetLotteryListLastIdResp               = pb.GetLotteryListLastIdResp
	GetLotteryPrizesListByUserIdReq        = pb.GetLotteryPrizesListByUserIdReq
	GetLotteryPrizesListByUserIdResp       = pb.GetLotteryPrizesListByUserIdResp
	GetLotteryStatisticReq                 = pb.GetLotteryStatisticReq
	GetLotteryStatisticResp                = pb.GetLotteryStatisticResp
	GetParticipationUserIdsByLotteryIdReq  = pb.GetParticipationUserIdsByLotteryIdReq
	GetParticipationUserIdsByLotteryIdResp = pb.GetParticipationUserIdsByLotteryIdResp
	GetPrizeByIdReq                        = pb.GetPrizeByIdReq
	GetPrizeByIdResp                       = pb.GetPrizeByIdResp
	GetPrizeListByLotteryIdReq             = pb.GetPrizeListByLotteryIdReq
	GetPrizeListByLotteryIdResp            = pb.GetPrizeListByLotteryIdResp
	GetSelectedLotteryStatisticReq         = pb.GetSelectedLotteryStatisticReq
	GetSelectedLotteryStatisticResp        = pb.GetSelectedLotteryStatisticResp
	GetWonListByLotteryIdReq               = pb.GetWonListByLotteryIdReq
	GetWonListByLotteryIdResp              = pb.GetWonListByLotteryIdResp
	GetWonListCountReq                     = pb.GetWonListCountReq
	GetWonListCountResp                    = pb.GetWonListCountResp
	GetWonListReq                          = pb.GetWonListReq
	GetWonListResp                         = pb.GetWonListResp
	Lottery                                = pb.Lottery
	LotteryDetailReq                       = pb.LotteryDetailReq
	LotteryDetailResp                      = pb.LotteryDetailResp
	LotteryParticipation                   = pb.LotteryParticipation
	LotteryPrizes                          = pb.LotteryPrizes
	LotterySponsorReq                      = pb.LotterySponsorReq
	LotterySponsorResp                     = pb.LotterySponsorResp
	Prize                                  = pb.Prize
	PublishLotteryReq                      = pb.PublishLotteryReq
	PublishLotteryResp                     = pb.PublishLotteryResp
	SearchLotteryParticipationReq          = pb.SearchLotteryParticipationReq
	SearchLotteryParticipationResp         = pb.SearchLotteryParticipationResp
	SearchLotteryReq                       = pb.SearchLotteryReq
	SearchLotteryResp                      = pb.SearchLotteryResp
	SearchPrizeReq                         = pb.SearchPrizeReq
	SearchPrizeResp                        = pb.SearchPrizeResp
	SetIsSelectedLotteryReq                = pb.SetIsSelectedLotteryReq
	SetIsSelectedLotteryResp               = pb.SetIsSelectedLotteryResp
	UpdateLotteryReq                       = pb.UpdateLotteryReq
	UpdateLotteryResp                      = pb.UpdateLotteryResp
	UpdatePrizeReq                         = pb.UpdatePrizeReq
	UpdatePrizeResp                        = pb.UpdatePrizeResp
	UserInfo                               = pb.UserInfo
	WonList                                = pb.WonList
	WonList2                               = pb.WonList2

	LotteryZrpcClient interface {
		// -------------------------抽奖表-------------------------
		AddLottery(ctx context.Context, in *AddLotteryReq, opts ...grpc.CallOption) (*AddLotteryResp, error)
		UpdateLottery(ctx context.Context, in *UpdateLotteryReq, opts ...grpc.CallOption) (*UpdateLotteryResp, error)
		DelLottery(ctx context.Context, in *DelLotteryReq, opts ...grpc.CallOption) (*DelLotteryResp, error)
		GetLotteryById(ctx context.Context, in *GetLotteryByIdReq, opts ...grpc.CallOption) (*GetLotteryByIdResp, error)
		SearchLottery(ctx context.Context, in *SearchLotteryReq, opts ...grpc.CallOption) (*SearchLotteryResp, error)
		SetIsSelectedLottery(ctx context.Context, in *SetIsSelectedLotteryReq, opts ...grpc.CallOption) (*SetIsSelectedLotteryResp, error)
		LotteryDetail(ctx context.Context, in *LotteryDetailReq, opts ...grpc.CallOption) (*LotteryDetailResp, error)
		LotterySponsor(ctx context.Context, in *LotterySponsorReq, opts ...grpc.CallOption) (*LotterySponsorResp, error)
		AnnounceLottery(ctx context.Context, in *AnnounceLotteryReq, opts ...grpc.CallOption) (*AnnounceLotteryResp, error)
		CheckUserCreateLottery(ctx context.Context, in *CheckUserCreateLotteryReq, opts ...grpc.CallOption) (*CheckUserCreateLotteryResp, error)
		CheckUserCreateLotteryAndToday(ctx context.Context, in *CheckUserCreateLotteryAndTodayReq, opts ...grpc.CallOption) (*CheckUserCreateLotteryAndTodayResp, error)
		CheckUserCreatedLotteryAndThisWeek(ctx context.Context, in *CheckUserCreatedLotteryAndThisWeekReq, opts ...grpc.CallOption) (*CheckUserCreatedLotteryAndThisWeekResp, error)
		GetLotteryListAfterLogin(ctx context.Context, in *GetLotteryListAfterLoginReq, opts ...grpc.CallOption) (*GetLotteryListAfterLoginResp, error)
		GetLotteryStatistic(ctx context.Context, in *GetLotteryStatisticReq, opts ...grpc.CallOption) (*GetLotteryStatisticResp, error)
		GetLotteryListLastId(ctx context.Context, in *GetLotteryListLastIdReq, opts ...grpc.CallOption) (*GetLotteryListLastIdResp, error)
		PublishLottery(ctx context.Context, in *PublishLotteryReq, opts ...grpc.CallOption) (*PublishLotteryResp, error)
		GetLotteryPrizesListByUserId(ctx context.Context, in *GetLotteryPrizesListByUserIdReq, opts ...grpc.CallOption) (*GetLotteryPrizesListByUserIdResp, error)
		// -------------------------奖品表-----------------------------------------------
		AddPrize(ctx context.Context, in *AddPrizeReq, opts ...grpc.CallOption) (*AddPrizeResp, error)
		UpdatePrize(ctx context.Context, in *UpdatePrizeReq, opts ...grpc.CallOption) (*UpdatePrizeResp, error)
		DelPrize(ctx context.Context, in *DelPrizeReq, opts ...grpc.CallOption) (*DelPrizeResp, error)
		GetPrizeById(ctx context.Context, in *GetPrizeByIdReq, opts ...grpc.CallOption) (*GetPrizeByIdResp, error)
		SearchPrize(ctx context.Context, in *SearchPrizeReq, opts ...grpc.CallOption) (*SearchPrizeResp, error)
		GetPrizeListByLotteryId(ctx context.Context, in *GetPrizeListByLotteryIdReq, opts ...grpc.CallOption) (*GetPrizeListByLotteryIdResp, error)
		// -------------------------参与抽奖表-------------------------
		AddLotteryParticipation(ctx context.Context, in *AddLotteryParticipationReq, opts ...grpc.CallOption) (*AddLotteryParticipationResp, error)
		SearchLotteryParticipation(ctx context.Context, in *SearchLotteryParticipationReq, opts ...grpc.CallOption) (*SearchLotteryParticipationResp, error)
		GetParticipationUserIdsByLotteryId(ctx context.Context, in *GetParticipationUserIdsByLotteryIdReq, opts ...grpc.CallOption) (*GetParticipationUserIdsByLotteryIdResp, error)
		CheckIsParticipated(ctx context.Context, in *CheckIsParticipatedReq, opts ...grpc.CallOption) (*CheckIsParticipatedResp, error)
		GetSelectedLotteryStatistic(ctx context.Context, in *GetSelectedLotteryStatisticReq, opts ...grpc.CallOption) (*GetSelectedLotteryStatisticResp, error)
		CheckSelectedLotteryParticipated(ctx context.Context, in *CheckSelectedLotteryParticipatedReq, opts ...grpc.CallOption) (*CheckSelectedLotteryParticipatedResp, error)
		CheckUserIsWon(ctx context.Context, in *CheckUserIsWonReq, opts ...grpc.CallOption) (*CheckUserIsWonResp, error)
		GetWonList(ctx context.Context, in *GetWonListReq, opts ...grpc.CallOption) (*GetWonListResp, error)
		GetWonListCount(ctx context.Context, in *GetWonListCountReq, opts ...grpc.CallOption) (*GetWonListCountResp, error)
		GetWonListByLotteryId(ctx context.Context, in *GetWonListByLotteryIdReq, opts ...grpc.CallOption) (*GetWonListByLotteryIdResp, error)
		// -----------------------完成打卡任务-----------------------
		AddClockTaskRecord(ctx context.Context, in *AddClockTaskRecordReq, opts ...grpc.CallOption) (*AddClockTaskRecordResp, error)
	}

	defaultLotteryZrpcClient struct {
		cli zrpc.Client
	}
)

func NewLotteryZrpcClient(cli zrpc.Client) LotteryZrpcClient {
	return &defaultLotteryZrpcClient{
		cli: cli,
	}
}

// -------------------------抽奖表-------------------------
func (m *defaultLotteryZrpcClient) AddLottery(ctx context.Context, in *AddLotteryReq, opts ...grpc.CallOption) (*AddLotteryResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.AddLottery(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) UpdateLottery(ctx context.Context, in *UpdateLotteryReq, opts ...grpc.CallOption) (*UpdateLotteryResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.UpdateLottery(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) DelLottery(ctx context.Context, in *DelLotteryReq, opts ...grpc.CallOption) (*DelLotteryResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.DelLottery(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetLotteryById(ctx context.Context, in *GetLotteryByIdReq, opts ...grpc.CallOption) (*GetLotteryByIdResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetLotteryById(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) SearchLottery(ctx context.Context, in *SearchLotteryReq, opts ...grpc.CallOption) (*SearchLotteryResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.SearchLottery(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) SetIsSelectedLottery(ctx context.Context, in *SetIsSelectedLotteryReq, opts ...grpc.CallOption) (*SetIsSelectedLotteryResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.SetIsSelectedLottery(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) LotteryDetail(ctx context.Context, in *LotteryDetailReq, opts ...grpc.CallOption) (*LotteryDetailResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.LotteryDetail(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) LotterySponsor(ctx context.Context, in *LotterySponsorReq, opts ...grpc.CallOption) (*LotterySponsorResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.LotterySponsor(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) AnnounceLottery(ctx context.Context, in *AnnounceLotteryReq, opts ...grpc.CallOption) (*AnnounceLotteryResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.AnnounceLottery(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) CheckUserCreateLottery(ctx context.Context, in *CheckUserCreateLotteryReq, opts ...grpc.CallOption) (*CheckUserCreateLotteryResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.CheckUserCreateLottery(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) CheckUserCreateLotteryAndToday(ctx context.Context, in *CheckUserCreateLotteryAndTodayReq, opts ...grpc.CallOption) (*CheckUserCreateLotteryAndTodayResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.CheckUserCreateLotteryAndToday(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) CheckUserCreatedLotteryAndThisWeek(ctx context.Context, in *CheckUserCreatedLotteryAndThisWeekReq, opts ...grpc.CallOption) (*CheckUserCreatedLotteryAndThisWeekResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.CheckUserCreatedLotteryAndThisWeek(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetLotteryListAfterLogin(ctx context.Context, in *GetLotteryListAfterLoginReq, opts ...grpc.CallOption) (*GetLotteryListAfterLoginResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetLotteryListAfterLogin(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetLotteryStatistic(ctx context.Context, in *GetLotteryStatisticReq, opts ...grpc.CallOption) (*GetLotteryStatisticResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetLotteryStatistic(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetLotteryListLastId(ctx context.Context, in *GetLotteryListLastIdReq, opts ...grpc.CallOption) (*GetLotteryListLastIdResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetLotteryListLastId(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) PublishLottery(ctx context.Context, in *PublishLotteryReq, opts ...grpc.CallOption) (*PublishLotteryResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.PublishLottery(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetLotteryPrizesListByUserId(ctx context.Context, in *GetLotteryPrizesListByUserIdReq, opts ...grpc.CallOption) (*GetLotteryPrizesListByUserIdResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetLotteryPrizesListByUserId(ctx, in, opts...)
}

// -------------------------奖品表-----------------------------------------------
func (m *defaultLotteryZrpcClient) AddPrize(ctx context.Context, in *AddPrizeReq, opts ...grpc.CallOption) (*AddPrizeResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.AddPrize(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) UpdatePrize(ctx context.Context, in *UpdatePrizeReq, opts ...grpc.CallOption) (*UpdatePrizeResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.UpdatePrize(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) DelPrize(ctx context.Context, in *DelPrizeReq, opts ...grpc.CallOption) (*DelPrizeResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.DelPrize(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetPrizeById(ctx context.Context, in *GetPrizeByIdReq, opts ...grpc.CallOption) (*GetPrizeByIdResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetPrizeById(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) SearchPrize(ctx context.Context, in *SearchPrizeReq, opts ...grpc.CallOption) (*SearchPrizeResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.SearchPrize(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetPrizeListByLotteryId(ctx context.Context, in *GetPrizeListByLotteryIdReq, opts ...grpc.CallOption) (*GetPrizeListByLotteryIdResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetPrizeListByLotteryId(ctx, in, opts...)
}

// -------------------------参与抽奖表-------------------------
func (m *defaultLotteryZrpcClient) AddLotteryParticipation(ctx context.Context, in *AddLotteryParticipationReq, opts ...grpc.CallOption) (*AddLotteryParticipationResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.AddLotteryParticipation(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) SearchLotteryParticipation(ctx context.Context, in *SearchLotteryParticipationReq, opts ...grpc.CallOption) (*SearchLotteryParticipationResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.SearchLotteryParticipation(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetParticipationUserIdsByLotteryId(ctx context.Context, in *GetParticipationUserIdsByLotteryIdReq, opts ...grpc.CallOption) (*GetParticipationUserIdsByLotteryIdResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetParticipationUserIdsByLotteryId(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) CheckIsParticipated(ctx context.Context, in *CheckIsParticipatedReq, opts ...grpc.CallOption) (*CheckIsParticipatedResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.CheckIsParticipated(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetSelectedLotteryStatistic(ctx context.Context, in *GetSelectedLotteryStatisticReq, opts ...grpc.CallOption) (*GetSelectedLotteryStatisticResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetSelectedLotteryStatistic(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) CheckSelectedLotteryParticipated(ctx context.Context, in *CheckSelectedLotteryParticipatedReq, opts ...grpc.CallOption) (*CheckSelectedLotteryParticipatedResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.CheckSelectedLotteryParticipated(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) CheckUserIsWon(ctx context.Context, in *CheckUserIsWonReq, opts ...grpc.CallOption) (*CheckUserIsWonResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.CheckUserIsWon(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetWonList(ctx context.Context, in *GetWonListReq, opts ...grpc.CallOption) (*GetWonListResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetWonList(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetWonListCount(ctx context.Context, in *GetWonListCountReq, opts ...grpc.CallOption) (*GetWonListCountResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetWonListCount(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetWonListByLotteryId(ctx context.Context, in *GetWonListByLotteryIdReq, opts ...grpc.CallOption) (*GetWonListByLotteryIdResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetWonListByLotteryId(ctx, in, opts...)
}

// -----------------------完成打卡任务-----------------------
func (m *defaultLotteryZrpcClient) AddClockTaskRecord(ctx context.Context, in *AddClockTaskRecordReq, opts ...grpc.CallOption) (*AddClockTaskRecordResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.AddClockTaskRecord(ctx, in, opts...)
}