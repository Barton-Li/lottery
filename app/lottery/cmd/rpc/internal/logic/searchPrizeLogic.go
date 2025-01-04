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

type SearchPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPrizeLogic {
	return &SearchPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPrizeLogic) SearchPrize(in *pb.SearchPrizeReq) (*pb.SearchPrizeResp, error) {
	prizes, err := l.svcCtx.PrizeModel.FindPageByLotteryId(l.ctx, in.LotteryId, in.Page, in.Limit)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "SearchPrizeLogic SearchPrize DB err , err : %v , in : %+v", err, in)
	}
	var resp []*pb.Prize
	if len(prizes) > 0 {
		for _, prize := range prizes {
			var pbPrize *pb.Prize
			_ = copier.Copy(&pbPrize, prize)
			resp = append(resp, pbPrize)
		}
	}
	return &pb.SearchPrizeResp{
		Prize: resp,
	}, nil
}
