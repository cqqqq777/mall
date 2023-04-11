// Code generated by Kitex v0.5.1. DO NOT EDIT.

package orderservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	order "github.com/cqqqq777/go-kitex-mall/shared/kitex_gen/order"
)

func serviceInfo() *kitex.ServiceInfo {
	return orderServiceServiceInfo
}

var orderServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "OrderService"
	handlerType := (*order.OrderService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateOrder": kitex.NewMethodInfo(createOrderHandler, newOrderServiceCreateOrderArgs, newOrderServiceCreateOrderResult, false),
		"UpdateOrder": kitex.NewMethodInfo(updateOrderHandler, newOrderServiceUpdateOrderArgs, newOrderServiceUpdateOrderResult, false),
		"OrderList":   kitex.NewMethodInfo(orderListHandler, newOrderServiceOrderListArgs, newOrderServiceOrderListResult, false),
		"GetOrder":    kitex.NewMethodInfo(getOrderHandler, newOrderServiceGetOrderArgs, newOrderServiceGetOrderResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "order",
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

func createOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceCreateOrderArgs)
	realResult := result.(*order.OrderServiceCreateOrderResult)
	success, err := handler.(order.OrderService).CreateOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceCreateOrderArgs() interface{} {
	return order.NewOrderServiceCreateOrderArgs()
}

func newOrderServiceCreateOrderResult() interface{} {
	return order.NewOrderServiceCreateOrderResult()
}

func updateOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceUpdateOrderArgs)
	realResult := result.(*order.OrderServiceUpdateOrderResult)
	success, err := handler.(order.OrderService).UpdateOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceUpdateOrderArgs() interface{} {
	return order.NewOrderServiceUpdateOrderArgs()
}

func newOrderServiceUpdateOrderResult() interface{} {
	return order.NewOrderServiceUpdateOrderResult()
}

func orderListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceOrderListArgs)
	realResult := result.(*order.OrderServiceOrderListResult)
	success, err := handler.(order.OrderService).OrderList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceOrderListArgs() interface{} {
	return order.NewOrderServiceOrderListArgs()
}

func newOrderServiceOrderListResult() interface{} {
	return order.NewOrderServiceOrderListResult()
}

func getOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceGetOrderArgs)
	realResult := result.(*order.OrderServiceGetOrderResult)
	success, err := handler.(order.OrderService).GetOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceGetOrderArgs() interface{} {
	return order.NewOrderServiceGetOrderArgs()
}

func newOrderServiceGetOrderResult() interface{} {
	return order.NewOrderServiceGetOrderResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateOrder(ctx context.Context, req *order.MallCreateOrderRequest) (r *order.MallCreateOrderResponse, err error) {
	var _args order.OrderServiceCreateOrderArgs
	_args.Req = req
	var _result order.OrderServiceCreateOrderResult
	if err = p.c.Call(ctx, "CreateOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateOrder(ctx context.Context, req *order.MallUpdateOrderRequest) (r *order.MallUpdateOrderResponse, err error) {
	var _args order.OrderServiceUpdateOrderArgs
	_args.Req = req
	var _result order.OrderServiceUpdateOrderResult
	if err = p.c.Call(ctx, "UpdateOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) OrderList(ctx context.Context, req *order.MallOrderListRequset) (r *order.MallOrderListResponse, err error) {
	var _args order.OrderServiceOrderListArgs
	_args.Req = req
	var _result order.OrderServiceOrderListResult
	if err = p.c.Call(ctx, "OrderList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetOrder(ctx context.Context, req *order.MallGetOrderRequest) (r *order.MallGetOrderResponse, err error) {
	var _args order.OrderServiceGetOrderArgs
	_args.Req = req
	var _result order.OrderServiceGetOrderResult
	if err = p.c.Call(ctx, "GetOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
