// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	shop "lottery/app/shop/cmd/api/internal/handler/shop"
	"lottery/app/shop/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/goodsInfo/getGoodsById",
				Handler: shop.GetGoodsByIdHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/shop/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/goodsInfo/getGoodsList",
				Handler: shop.GetGoodsListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/shop/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/goodsInfo/syncPddGoods",
				Handler: shop.SyncPddGoodsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/shop/v1"),
	)
}