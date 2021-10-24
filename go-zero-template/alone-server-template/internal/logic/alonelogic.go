package logic

import (
	"context"

	"alone-server-template/internal/svc"
	"alone-server-template/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AloneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAloneLogic(ctx context.Context, svcCtx *svc.ServiceContext) AloneLogic {
	return AloneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AloneLogic) Alone(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line
	return &types.Response{
		Message: "Hello, I is alone server!",
	}, nil
}
