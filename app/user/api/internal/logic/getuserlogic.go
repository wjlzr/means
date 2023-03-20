package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
	"means/app/user/api/internal/svc"
	"means/app/user/api/internal/types"
	"means/app/user/rpc/user"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUser
//  @Description:
//  @receiver l
//  @param req
//  @return resp
//  @return err
func (l *GetUserLogic) GetUser(req *types.GetUserReq) (*types.User, error) {

	userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var resp types.User
	_ = copier.Copy(&resp, userResp)

	return &resp, nil
}
