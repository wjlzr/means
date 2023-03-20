package logic

import (
	"context"
	"github.com/pkg/errors"
	"means/app/user/model"
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

// GetUser
// @Description 获取用户信息
// @Author bryce
// @Date 2023-03-17 13:47:26
func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.UserInfo, error) {

	userRes, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.INVALID_ERROR), "Logic GetUser FindOne ERR=====>%s", err.Error())
		} else {
			return nil, xerr.NewErrCode(xerr.DB_ERROR)
		}
	}

	return &user.UserInfo{
		Id:       userRes.Id,
		Nickname: userRes.Nickname,
		Avatar:   userRes.Avatar,
		Phone:    userRes.Phone,
	}, nil
}
