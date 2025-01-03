package address

import (
	"context"

	"lottery/app/usercenter/cmd/api/internal/svc"
	"lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 识别并转换收货地址
func NewConvertAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertAddressLogic {
	return &ConvertAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertAddressLogic) ConvertAddress(req *types.ConvertAddressReq) (resp *types.ConvertAddressResp, err error) {
	// todo: add your logic here and delete this line

	return
}
