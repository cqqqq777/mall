// Code generated by Kitex v0.5.1. DO NOT EDIT.

package merchantservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	merchant "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/merchant"
)

func serviceInfo() *kitex.ServiceInfo {
	return merchantServiceServiceInfo
}

var merchantServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MerchantService"
	handlerType := (*merchant.MerchantService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register": kitex.NewMethodInfo(registerHandler, newMerchantServiceRegisterArgs, newMerchantServiceRegisterResult, false),
		"Login":    kitex.NewMethodInfo(loginHandler, newMerchantServiceLoginArgs, newMerchantServiceLoginResult, false),
		"GetInfo":  kitex.NewMethodInfo(getInfoHandler, newMerchantServiceGetInfoArgs, newMerchantServiceGetInfoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "merchant",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*merchant.MerchantServiceRegisterArgs)
	realResult := result.(*merchant.MerchantServiceRegisterResult)
	success, err := handler.(merchant.MerchantService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMerchantServiceRegisterArgs() interface{} {
	return merchant.NewMerchantServiceRegisterArgs()
}

func newMerchantServiceRegisterResult() interface{} {
	return merchant.NewMerchantServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*merchant.MerchantServiceLoginArgs)
	realResult := result.(*merchant.MerchantServiceLoginResult)
	success, err := handler.(merchant.MerchantService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMerchantServiceLoginArgs() interface{} {
	return merchant.NewMerchantServiceLoginArgs()
}

func newMerchantServiceLoginResult() interface{} {
	return merchant.NewMerchantServiceLoginResult()
}

func getInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*merchant.MerchantServiceGetInfoArgs)
	realResult := result.(*merchant.MerchantServiceGetInfoResult)
	success, err := handler.(merchant.MerchantService).GetInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMerchantServiceGetInfoArgs() interface{} {
	return merchant.NewMerchantServiceGetInfoArgs()
}

func newMerchantServiceGetInfoResult() interface{} {
	return merchant.NewMerchantServiceGetInfoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *merchant.MallMerchantRegisterRequest) (r *merchant.MallMerchantRegisterResponse, err error) {
	var _args merchant.MerchantServiceRegisterArgs
	_args.Req = req
	var _result merchant.MerchantServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *merchant.MallMerchantLoginRequest) (r *merchant.MallMerchantLoginResponse, err error) {
	var _args merchant.MerchantServiceLoginArgs
	_args.Req = req
	var _result merchant.MerchantServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetInfo(ctx context.Context, req *merchant.MallMerchantGetInfoRequest) (r *merchant.MallMerchantGetInfoResponse, err error) {
	var _args merchant.MerchantServiceGetInfoArgs
	_args.Req = req
	var _result merchant.MerchantServiceGetInfoResult
	if err = p.c.Call(ctx, "GetInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
