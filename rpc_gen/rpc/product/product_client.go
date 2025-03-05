package product

import (
	"context"

	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"

	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() productcatalogservice.Client
	Service() string
	ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error)
	GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error)
	InsertProducts(ctx context.Context, Req *product.InsertProductsReq, callOptions ...callopt.Option) (r *product.Empty, err error)
	DeleteProducts(ctx context.Context, Req *product.DeleteProductsReq, callOptions ...callopt.Option) (r *product.Empty, err error)
	UpdateProducts(ctx context.Context, Req *product.UpdateProductsReq, callOptions ...callopt.Option) (r *product.Empty, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productcatalogservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productcatalogservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productcatalogservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error) {
	return c.kitexClient.ListProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	return c.kitexClient.GetProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error) {
	return c.kitexClient.SearchProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) InsertProducts(ctx context.Context, Req *product.InsertProductsReq, callOptions ...callopt.Option) (r *product.Empty, err error) {
	return c.kitexClient.InsertProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteProducts(ctx context.Context, Req *product.DeleteProductsReq, callOptions ...callopt.Option) (r *product.Empty, err error) {
	return c.kitexClient.DeleteProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateProducts(ctx context.Context, Req *product.UpdateProductsReq, callOptions ...callopt.Option) (r *product.Empty, err error) {
	return c.kitexClient.UpdateProducts(ctx, Req, callOptions...)
}
