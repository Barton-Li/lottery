package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"lottery/app/usercenter/cmd/rpc/internal/svc"
	pb2 "lottery/app/usercenter/cmd/rpc/pb"
	"lottery/common/xerr"

	"lottery/app/usercenter/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserSponsorLogic {
	return &SearchUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchUserSponsor 根据用户ID查询用户推广信息。
// in 参数包括用户ID、页码和每页数量，用于分页查询。
// 返回用户推广信息列表和潜在的错误。
func (l *SearchUserSponsorLogic) SearchUserSponsor(in *pb2.SearchUserSponsorReq) (*pb2.SearchUserSponsorResp, error) {
    // 调用模型层方法，根据用户ID、页码和每页数量查询推广信息列表。
    list, err := l.svcCtx.UserSponsorModel.FindPageByUserId(l.ctx, in.UserId, in.Page, in.Limit)
    // 如果发生错误且不是未找到记录的错误，则包装错误信息并返回。
    if err!=nil&&err!=model.ErrNotFound {
        return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询推广信息失败err:%v,req:%v", err,in)
    }

    // 初始化响应列表。
    var resp []*pb2.UserSponsor
    // 遍历查询结果，将每条记录转换为响应对象并添加到响应列表中。
    for len(list)>0{
        for _,sponsor:=range list{
            // 创建一个新的UserSponsor对象以存储转换后的数据。
            var pbSponsor pb2.UserSponsor
            // 使用copier库复制数据，忽略错误处理因为此处转换逻辑简单且数据结构已知。
            _=copier.Copy(&pbSponsor,sponsor)
            // 将转换后的对象添加到响应列表中。
            resp = append(resp, &pbSponsor)
        }
    }

    // 返回包含用户推广信息列表的响应对象。
    return &pb2.SearchUserSponsorResp{
        UserSponsor: resp,
    }, nil
}

