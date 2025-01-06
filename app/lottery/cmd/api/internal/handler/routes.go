// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	lottery "lottery/app/lottery/cmd/api/internal/handler/lottery"
	"lottery/app/lottery/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 测试Validator
				Method:  http.MethodPost,
				Path:    "/lottery/TestValidator",
				Handler: lottery.TestValidatorHandler(serverCtx),
			},
			{
				// 增加概率类型
				Method:  http.MethodGet,
				Path:    "/lottery/chanceTypeList",
				Handler: lottery.ChanceTypeListHandler(serverCtx),
			},
			{
				// 打卡任务类型
				Method:  http.MethodGet,
				Path:    "/lottery/clockTaskTypeList",
				Handler: lottery.ClockTaskTypeListHandler(serverCtx),
			},
			{
				// 获取当前抽奖中奖者名单
				Method:  http.MethodPost,
				Path:    "/lottery/getLotteryWinnersList",
				Handler: lottery.GetLotteryWinlist2Handler(serverCtx),
			},
			{
				// 抽奖列表
				Method:  http.MethodGet,
				Path:    "/lottery/lotteryList",
				Handler: lottery.LotteryListHandler(serverCtx),
			},
			{
				// 抽奖人
				Method:  http.MethodPost,
				Path:    "/lottery/participations",
				Handler: lottery.SearchParticipationHandler(serverCtx),
			},
		},
		rest.WithPrefix("/lottery/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// CheckIsParticipated
				Method:  http.MethodPost,
				Path:    "/lottery/CheckIsParticipated",
				Handler: lottery.CheckIsParticipatedHandler(serverCtx),
			},
			{
				// 判断当前用户当前抽奖是否中奖
				Method:  http.MethodPost,
				Path:    "/lottery/checkIsWin",
				Handler: lottery.CheckIsWinHandler(serverCtx),
			},
			{
				// 完成打卡任务
				Method:  http.MethodPost,
				Path:    "/lottery/createClockTaskRecord",
				Handler: lottery.CreateClockTaskRecordHandler(serverCtx),
			},
			{
				// 发起抽奖
				Method:  http.MethodPost,
				Path:    "/lottery/createLottery",
				Handler: lottery.CreateLotteryHandler(serverCtx),
			},
			{
				// 获取当前用户全部/发起/中奖的抽奖列表
				Method:  http.MethodPost,
				Path:    "/lottery/getLotteryListByUserId",
				Handler: lottery.GetLotteryListByUserIdHandler(serverCtx),
			},
			{
				// 获取当前用户中奖列表
				Method:  http.MethodPost,
				Path:    "/lottery/getLotteryWinList",
				Handler: lottery.GetLotteryWinListHandler(serverCtx),
			},
			{
				// lottery detail
				Method:  http.MethodPost,
				Path:    "/lottery/lotteryDetail",
				Handler: lottery.LotteryDetailHandler(serverCtx),
			},
			{
				// 登录后获取抽奖列表
				Method:  http.MethodPost,
				Path:    "/lottery/lotteryListAfterLogin",
				Handler: lottery.LotteryListAfterLoginHandler(serverCtx),
			},
			{
				// 参与抽奖
				Method:  http.MethodPost,
				Path:    "/lottery/participation",
				Handler: lottery.AddLotteryParticipationHandler(serverCtx),
			},
			{
				// lottery setIsSelected
				Method:  http.MethodPost,
				Path:    "/lottery/setLotteryIsSelected",
				Handler: lottery.SetLotteryIsSelectedHandler(serverCtx),
			},
			{
				// 发布抽奖
				Method:  http.MethodPost,
				Path:    "/lottery/updateLottery",
				Handler: lottery.UpdateLotteryHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.jwtAuth.AccessSecret),
		rest.WithPrefix("/lottery/v1"),
	)
}