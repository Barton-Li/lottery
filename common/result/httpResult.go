package result

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"lottery/common/xerr"
	"net/http"
)

//http返回
// HttpResult 根据处理请求的结果，生成HTTP响应。
// 该函数接收一个HTTP请求对象、一个HTTP响应写入对象、一个响应体接口和一个错误对象。
// 如果没有错误，它会生成一个成功的HTTP响应；如果有错误，它会根据错误的类型生成一个适当的错误响应。
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
    // 当没有错误时，生成成功的HTTP响应。
	if err == nil {
		// 成功时，调用Success函数处理响应体，并使用httpx.WriteJson发送HTTP 200响应。
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
        // 定义默认的错误代码和消息。
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := "服务器内部错误,请稍后再试"

        // 获取错误的根本原因。
		causeErr := errors.Cause(err)
        // 检查错误是否为自定义的CodeError类型，如果是，则使用其错误代码和消息。
		if e, ok := causeErr.(*xerr.CodeError); ok {
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
            // 检查错误是否来自gRPC，如果是，并且是已知的错误代码，则使用相应的错误代码和消息。
			if gStatus, ok := status.FromError(err); ok {
				grpcCode := uint32(gStatus.Code())
				if xerr.IsCoedErr(grpcCode) {
					errCode = grpcCode
					errMsg = gStatus.Message()
				}

			}
		}
        // 记录错误日志。
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v 【API-ERR-END】", err)
        // 生成错误的HTTP响应，状态码为400。
		httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
	}
}


//授权http返回
func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := "服务器内部错误,请稍后再试"

		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*xerr.CodeError); ok {
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gStatus, ok := status.FromError(err); ok {
				grpcCode := uint32(gStatus.Code())
				if xerr.IsCoedErr(grpcCode) {
					errCode = grpcCode
					errMsg = gStatus.Message()
				}

			}
		}
		logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v 【API-ERR-END】", err)
		httpx.WriteJson(w, http.StatusUnauthorized, Error(errCode, errMsg))
	}
}

//参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.REUQEST_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.REUQEST_PARAM_ERROR, errMsg))

}
