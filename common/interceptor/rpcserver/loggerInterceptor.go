package rpcserver

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"means/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

// LoggerInterceptor
// @Des rpc service logger interceptor
// @Author bryce
// @Date 2023-03-17 13:58:27
func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %s", err.Error())
			//转成grpc err
			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
		} else {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
		}
	}
	return resp, err
}
