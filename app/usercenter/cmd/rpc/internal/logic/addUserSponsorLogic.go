package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"
	"lottery/app/usercenter/model"
	"lottery/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserSponsorLogic {
	return &AddUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------userSponsor-----------------------
// AddUserSponsorLogic.AddUserSponsor 添加用户赞助信息。
// 该方法接收一个AddUserSponsorReq请求对象，尝试将其中的数据转换为UserSponsor模型对象，并插入到数据库中。
// 如果插入成功，返回新插入记录的ID；如果失败，返回相应的错误。
func (l *AddUserSponsorLogic) AddUserSponsor(in *pb2.AddUserSponsorReq) (*pb2.AddUserSponsorResp, error) {
	// 创建一个新的UserSponsor对象，用于存储转换后的用户赞助信息。
	userSponsor := new(model.UserSponsor)
	// 使用copier库将请求对象in中的数据复制到userSponsor对象中。
	err := copier.Copy(userSponsor, in)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to copy data:%+v,err:%s", in, err)
	}
	// 调用UserSponsorModel的Insert方法，将userSponsor对象插入到数据库中。
	insertResult, err := l.svcCtx.UserSponsorModel.Insert(l.ctx, userSponsor)
	if err != nil {
		// 如果插入过程中出现错误，返回nil和错误信息。
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "fFailed to insert user sponsor:%+v,err:%s", insertResult, err)

	}
	// 获取插入记录的最后插入ID。
	lastInsertId, err := insertResult.LastInsertId()
	if err != nil {
		// 如果获取最后插入ID时出现错误，封装错误信息并返回。
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add sponsor fail %d,err:%s", lastInsertId, err)
	}

	// 返回添加用户赞助信息的成功响应，包含新插入记录的ID。
	return &pb2.AddUserSponsorResp{
		Id: lastInsertId,
	}, nil
}
