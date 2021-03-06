// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: asset.proto

/*
Package asset is a generated protocol buffer package.

It is generated from these files:
	asset.proto

It has these top-level messages:
	GetFileUploadURLRequest
	GetFileUploadURLResponse
	RegisterFileRequest
	RegisterFileResponse
	QueryUploadedDataRequest
	QueryUploadedDataResponse
	QueryUploadedData
	QueryUploadedRow
	RegisterRequest
	RegisterResponse
	QueryRequest
	QueryResponse
	QueryData
	QueryRow
	QueryAllAssetRequest
	QueryPara
	QueryAllAssetResponse
	ModifyRequest
	ModifyResponse
	GetFileUploadStatRequest
	GetFileUploadStatResponse
	GetDownLoadURLRequest
	GetDownLoadURLResponse
	QueryByIDRequest
	GetUserPurchaseAssetListRequest
	GetUserPurchaseAssetListResponse
	QueryPurchaseData
	QueryPurchaseRow
*/
package asset

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Asset service

type AssetClient interface {
	GetFileUploadURL(ctx context.Context, in *GetFileUploadURLRequest, opts ...client.CallOption) (*GetFileUploadURLResponse, error)
	GetFileUploadStat(ctx context.Context, in *GetFileUploadStatRequest, opts ...client.CallOption) (*GetFileUploadStatResponse, error)
	RegisterFile(ctx context.Context, in *RegisterFileRequest, opts ...client.CallOption) (*RegisterFileResponse, error)
	QueryUploadedData(ctx context.Context, in *QueryUploadedDataRequest, opts ...client.CallOption) (*QueryUploadedDataResponse, error)
	GetDownLoadURL(ctx context.Context, in *GetDownLoadURLRequest, opts ...client.CallOption) (*GetDownLoadURLResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error)
	QueryAllAsset(ctx context.Context, in *QueryAllAssetRequest, opts ...client.CallOption) (*QueryAllAssetResponse, error)
	Modify(ctx context.Context, in *ModifyRequest, opts ...client.CallOption) (*ModifyResponse, error)
	QueryByID(ctx context.Context, in *QueryByIDRequest, opts ...client.CallOption) (*QueryResponse, error)
	GetUserPurchaseAssetList(ctx context.Context, in *GetUserPurchaseAssetListRequest, opts ...client.CallOption) (*QueryResponse, error)
}

type assetClient struct {
	c           client.Client
	serviceName string
}

func NewAssetClient(serviceName string, c client.Client) AssetClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "asset"
	}
	return &assetClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *assetClient) GetFileUploadURL(ctx context.Context, in *GetFileUploadURLRequest, opts ...client.CallOption) (*GetFileUploadURLResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.GetFileUploadURL", in)
	out := new(GetFileUploadURLResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) GetFileUploadStat(ctx context.Context, in *GetFileUploadStatRequest, opts ...client.CallOption) (*GetFileUploadStatResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.GetFileUploadStat", in)
	out := new(GetFileUploadStatResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) RegisterFile(ctx context.Context, in *RegisterFileRequest, opts ...client.CallOption) (*RegisterFileResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.RegisterFile", in)
	out := new(RegisterFileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) QueryUploadedData(ctx context.Context, in *QueryUploadedDataRequest, opts ...client.CallOption) (*QueryUploadedDataResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.QueryUploadedData", in)
	out := new(QueryUploadedDataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) GetDownLoadURL(ctx context.Context, in *GetDownLoadURLRequest, opts ...client.CallOption) (*GetDownLoadURLResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.GetDownLoadURL", in)
	out := new(GetDownLoadURLResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.Query", in)
	out := new(QueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) QueryAllAsset(ctx context.Context, in *QueryAllAssetRequest, opts ...client.CallOption) (*QueryAllAssetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.QueryAllAsset", in)
	out := new(QueryAllAssetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) Modify(ctx context.Context, in *ModifyRequest, opts ...client.CallOption) (*ModifyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.Modify", in)
	out := new(ModifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) QueryByID(ctx context.Context, in *QueryByIDRequest, opts ...client.CallOption) (*QueryResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.QueryByID", in)
	out := new(QueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) GetUserPurchaseAssetList(ctx context.Context, in *GetUserPurchaseAssetListRequest, opts ...client.CallOption) (*QueryResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Asset.GetUserPurchaseAssetList", in)
	out := new(QueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Asset service

type AssetHandler interface {
	GetFileUploadURL(context.Context, *GetFileUploadURLRequest, *GetFileUploadURLResponse) error
	GetFileUploadStat(context.Context, *GetFileUploadStatRequest, *GetFileUploadStatResponse) error
	RegisterFile(context.Context, *RegisterFileRequest, *RegisterFileResponse) error
	QueryUploadedData(context.Context, *QueryUploadedDataRequest, *QueryUploadedDataResponse) error
	GetDownLoadURL(context.Context, *GetDownLoadURLRequest, *GetDownLoadURLResponse) error
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
	Query(context.Context, *QueryRequest, *QueryResponse) error
	QueryAllAsset(context.Context, *QueryAllAssetRequest, *QueryAllAssetResponse) error
	Modify(context.Context, *ModifyRequest, *ModifyResponse) error
	QueryByID(context.Context, *QueryByIDRequest, *QueryResponse) error
	GetUserPurchaseAssetList(context.Context, *GetUserPurchaseAssetListRequest, *QueryResponse) error
}

func RegisterAssetHandler(s server.Server, hdlr AssetHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Asset{hdlr}, opts...))
}

type Asset struct {
	AssetHandler
}

func (h *Asset) GetFileUploadURL(ctx context.Context, in *GetFileUploadURLRequest, out *GetFileUploadURLResponse) error {
	return h.AssetHandler.GetFileUploadURL(ctx, in, out)
}

func (h *Asset) GetFileUploadStat(ctx context.Context, in *GetFileUploadStatRequest, out *GetFileUploadStatResponse) error {
	return h.AssetHandler.GetFileUploadStat(ctx, in, out)
}

func (h *Asset) RegisterFile(ctx context.Context, in *RegisterFileRequest, out *RegisterFileResponse) error {
	return h.AssetHandler.RegisterFile(ctx, in, out)
}

func (h *Asset) QueryUploadedData(ctx context.Context, in *QueryUploadedDataRequest, out *QueryUploadedDataResponse) error {
	return h.AssetHandler.QueryUploadedData(ctx, in, out)
}

func (h *Asset) GetDownLoadURL(ctx context.Context, in *GetDownLoadURLRequest, out *GetDownLoadURLResponse) error {
	return h.AssetHandler.GetDownLoadURL(ctx, in, out)
}

func (h *Asset) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.AssetHandler.Register(ctx, in, out)
}

func (h *Asset) Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error {
	return h.AssetHandler.Query(ctx, in, out)
}

func (h *Asset) QueryAllAsset(ctx context.Context, in *QueryAllAssetRequest, out *QueryAllAssetResponse) error {
	return h.AssetHandler.QueryAllAsset(ctx, in, out)
}

func (h *Asset) Modify(ctx context.Context, in *ModifyRequest, out *ModifyResponse) error {
	return h.AssetHandler.Modify(ctx, in, out)
}

func (h *Asset) QueryByID(ctx context.Context, in *QueryByIDRequest, out *QueryResponse) error {
	return h.AssetHandler.QueryByID(ctx, in, out)
}

func (h *Asset) GetUserPurchaseAssetList(ctx context.Context, in *GetUserPurchaseAssetListRequest, out *QueryResponse) error {
	return h.AssetHandler.GetUserPurchaseAssetList(ctx, in, out)
}
