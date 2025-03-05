package product

import (
	"context"

	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (resp *product.ListProductsResp, err error) {
	resp, err = defaultClient.ListProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *product.SearchProductsReq, callOptions ...callopt.Option) (resp *product.SearchProductsResp, err error) {
	resp, err = defaultClient.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func InsertProducts(ctx context.Context, req *product.InsertProductsReq, callOptions ...callopt.Option) (resp *product.Empty, err error) {
	resp, err = defaultClient.InsertProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "InsertProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProducts(ctx context.Context, req *product.DeleteProductsReq, callOptions ...callopt.Option) (resp *product.Empty, err error) {
	resp, err = defaultClient.DeleteProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateProducts(ctx context.Context, req *product.UpdateProductsReq, callOptions ...callopt.Option) (resp *product.Empty, err error) {
	resp, err = defaultClient.UpdateProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
