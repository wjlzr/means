package logic

import (
	"context"

	"means/app/user/rpc/internal/svc"
	"means/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *user.AddUserRequest) (*user.UserInfo, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfo{}, nil
}
