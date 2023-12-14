package greetlogic

import (
	"context"

	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/example/rpc/hello/internal/svc"
	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/example/rpc/hello/pb/hello"

	"github.com/wuntsong-org/go-zero-plus/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *hello.HelloReq) (*hello.HelloResp, error) {
	// todo: add your logic here and delete this line

	return &hello.HelloResp{}, nil
}
