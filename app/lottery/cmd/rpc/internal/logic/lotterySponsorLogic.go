package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"lottery/app/usercenter/cmd/rpc/usercenter"

	"lottery/app/lottery/cmd/rpc/internal/svc"
	"lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotterySponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLotterySponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotterySponsorLogic {
	return &LotterySponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LotterySponsorLogic) LotterySponsor(in *pb.LotterySponsorReq) (*pb.LotterySponsorResp, error) {
	detail, err := l.svcCtx.UserCenterRpc.SponsorDetail(l.ctx, &usercenter.SponsorDetailReq{
		Id: in.SponsorId,
	})
	if err != nil || detail == nil {
		return nil, err
	}
	pbSponsorResp := new(pb.LotterySponsorResp)
	err = copier.Copy(pbSponsorResp, detail)
	if err != nil {
		return nil, err
	}
	return pbSponsorResp, nil
}
