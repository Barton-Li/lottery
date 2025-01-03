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

type AddUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserAddressLogic {
	return &AddUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户收货地址表-----------------------

// AddUserAddress AddUserAddressLogic.AddUserAddress 添加用户地址信息。
// 该方法接收一个AddUserAddressReq请求对象，将其中的数据复制到UserAddress模型实例中，
// 然后将该实例插入到数据库中，并返回插入记录的ID。
// 主要涉及的数据结构有：
// - in (*pb2.AddUserAddressReq): 请求参数，包含用户地址的相关信息。
// 返回值为新插入用户地址的ID及可能的错误。
func (l *AddUserAddressLogic) AddUserAddress(in *pb2.AddUserAddressReq) (*pb2.AddUserAddressResp, error) {
	// 创建一个新的UserAddress实例，用于存储用户地址信息。
	userAddress := new(model.UserAddress)

	// 将请求参数in中的数据复制到userAddress中，以简化数据处理过程。
	err := copier.Copy(userAddress, in)
	if err != nil {
		// 如果数据复制过程中出现错误，则返回错误信息。
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copier:%+v,err:%v", in, err)
	}

	// 调用UserAddressModel的Insert方法，将userAddress实例插入到数据库中。
	insertResult, err := l.svcCtx.UserAddressModel.Insert(l.ctx, userAddress)
	if err != nil {
		// 如果数据库插入操作失败，则返回错误信息。
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add address db user_address Insert err:%v, address:%+v", err, userAddress)
	}

	// 获取最后一次插入操作生成的ID。
	lastId, err := insertResult.LastInsertId()
	if err != nil {
		// 如果获取LastInsertId失败，则返回错误信息。
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add address db user_address insertResult.LastInsertId err:%v, address:%+v", err, userAddress)
	}

	// 返回新插入用户地址的ID。
	return &pb2.AddUserAddressResp{
		Id: lastId,
	}, nil
}
