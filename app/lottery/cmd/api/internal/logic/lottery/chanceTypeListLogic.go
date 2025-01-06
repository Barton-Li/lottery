package lottery

import (
	"context"
	"fmt"
	"lottery/common/constants"

	"lottery/app/lottery/cmd/api/internal/svc"
	"lottery/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChanceTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChanceTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChanceTypeListLogic {
	return &ChanceTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ChanceTypeList 生成机会类型列表。
//
// 参数：
// - req: 请求参数，用于未来扩展功能时传递条件或参数。当前实现中未使用。
//
// 返回值：
// - resp: 包含生成的机会类型列表的响应对象。
// - err: 如果发生错误，则返回错误信息；否则返回 nil。
//
// 该函数首先创建一个随机类型的机会，并将其添加到列表中，
// 然后循环生成10个指定类型的机会，并依次添加到列表中。
func (l *ChanceTypeListLogic) ChanceTypeList(req *types.ChanceTypeListReq) (resp *types.ChanceTypeListResp, err error) {
	// 初始化机会类型列表
	var ChanceTypeList []types.ChanceType

	// 创建并初始化一个机会类型对象
	changeType := types.ChanceType{}

	// 设置随机类型的机会，并添加到列表中
	changeType.Type = constants.Random
	changeType.Text = constants.RandomText
	ChanceTypeList = append(ChanceTypeList, changeType)

	// 循环生成10个指定类型的机会，并添加到列表中
	for i := 1; i <= 10; i++ {
		changeType.Type = constants.Appoint
		changeType.Text = fmt.Sprintf(constants.AppointText, i)
		ChanceTypeList = append(ChanceTypeList, changeType)
	}

	// 返回包含机会类型列表的响应对象
	return &types.ChanceTypeListResp{List: ChanceTypeList}, nil
}
