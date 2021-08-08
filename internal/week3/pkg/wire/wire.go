package wire

import (
	// 标准包

	// 第三方包
	"github.com/google/wire"

	// 内部包
	"gotraining3/internal/week3/biz"
	"gotraining3/internal/week3/data"
	"gotraining3/internal/week3/pkg"
	"gotraining3/internal/week3/service"
)

func InitService() *service.PictureService {
	wire.Build(service.NewPictrueService, biz.NewPictrueUsecase, data.NewPictureRepo, pkg.NewDB)
	return &service.PictureService{}
}
