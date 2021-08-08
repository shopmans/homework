package service

import (
	// 标准包
	"context"
	// 第三方包
	// 内部包
	"gotraining3/api"
	"gotraining3/internal/week3/biz"
)

type PictureService struct {
	puc *biz.PictrueUsecase
}

func NewPictrueService(puc *biz.PictrueUsecase) *PictureService {
	return &PictureService{puc: puc}
}

func (svr *PictureService) ModifyPicture(ctx context.Context, r *api.PictureRequest) (*api.PictureReply, error) {
	pdo := biz.Picture{}
	pdo.Name = r.Name

	err := svr.puc.Modify(&pdo)
	return &api.PictureReply{Message: "OK"}, err
}
