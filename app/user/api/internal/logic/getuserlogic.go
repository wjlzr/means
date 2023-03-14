package logic

import (
	"context"

	"means/app/user/api/internal/svc"
	"means/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.UserReply, err error) {

	//loginResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &usercenter.LoginReq{
	//	AuthType: model.UserAuthTypeSystem,
	//	AuthKey:  req.Mobile,
	//	Password: req.Password,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//var resp types.LoginResp
	//_ = copier.Copy(&resp, loginResp)
	//
	//return &resp, nil
	return
}
