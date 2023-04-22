// Code generated by Kitex v0.5.1. DO NOT EDIT.

package cartservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	cart "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/cart"
)

func serviceInfo() *kitex.ServiceInfo {
	return cartServiceServiceInfo
}

var cartServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CartService"
	handlerType := (*cart.CartService)(nil)
	methods := map[string]kitex.MethodInfo{
		"AddProductToCart": kitex.NewMethodInfo(addProductToCartHandler, newCartServiceAddProductToCartArgs, newCartServiceAddProductToCartResult, false),
		"GetCart":          kitex.NewMethodInfo(getCartHandler, newCartServiceGetCartArgs, newCartServiceGetCartResult, false),
		"DelCartProduct":   kitex.NewMethodInfo(delCartProductHandler, newCartServiceDelCartProductArgs, newCartServiceDelCartProductResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "cart",
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

func addProductToCartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*cart.CartServiceAddProductToCartArgs)
	realResult := result.(*cart.CartServiceAddProductToCartResult)
	success, err := handler.(cart.CartService).AddProductToCart(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCartServiceAddProductToCartArgs() interface{} {
	return cart.NewCartServiceAddProductToCartArgs()
}

func newCartServiceAddProductToCartResult() interface{} {
	return cart.NewCartServiceAddProductToCartResult()
}

func getCartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*cart.CartServiceGetCartArgs)
	realResult := result.(*cart.CartServiceGetCartResult)
	success, err := handler.(cart.CartService).GetCart(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCartServiceGetCartArgs() interface{} {
	return cart.NewCartServiceGetCartArgs()
}

func newCartServiceGetCartResult() interface{} {
	return cart.NewCartServiceGetCartResult()
}

func delCartProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*cart.CartServiceDelCartProductArgs)
	realResult := result.(*cart.CartServiceDelCartProductResult)
	success, err := handler.(cart.CartService).DelCartProduct(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCartServiceDelCartProductArgs() interface{} {
	return cart.NewCartServiceDelCartProductArgs()
}

func newCartServiceDelCartProductResult() interface{} {
	return cart.NewCartServiceDelCartProductResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AddProductToCart(ctx context.Context, req *cart.MallAddProductToCartRequest) (r *cart.MallAddProductToCartResponse, err error) {
	var _args cart.CartServiceAddProductToCartArgs
	_args.Req = req
	var _result cart.CartServiceAddProductToCartResult
	if err = p.c.Call(ctx, "AddProductToCart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCart(ctx context.Context, req *cart.MallGetCartRequest) (r *cart.MallGetCartResponse, err error) {
	var _args cart.CartServiceGetCartArgs
	_args.Req = req
	var _result cart.CartServiceGetCartResult
	if err = p.c.Call(ctx, "GetCart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DelCartProduct(ctx context.Context, req *cart.MallDelCartProductRequest) (r *cart.MallDelCartProductResponse, err error) {
	var _args cart.CartServiceDelCartProductArgs
	_args.Req = req
	var _result cart.CartServiceDelCartProductResult
	if err = p.c.Call(ctx, "DelCartProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
