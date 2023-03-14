package logic

import (
	"context"
	"means/common/xerr"

	"means/app/user/rpc/internal/svc"
	"means/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.UserInfo, error) {

	userRes, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Rpc login GetUser FindOne ERR:", err.Error())
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}

	return &user.UserInfo{
		Id:       userRes.Id,
		Nickname: userRes.Nickname,
		Avatar:   userRes.Avatar,
		Phone:    userRes.Phone,
	}, nil
}
