package biz

// Domain OBject 对象
type Picture struct {
	Name string
}

// repo 接口
type PictureRepo interface {
	Modify(*Picture) error
}

// PictrueUsecase
type PictrueUsecase struct {
	repo PictureRepo
}

func NewPictrueUsecase(repo PictureRepo) *PictrueUsecase {
	return &PictrueUsecase{repo: repo}
}

func (pu *PictrueUsecase) Modify(p *Picture) error {
	return pu.repo.Modify(p)
}
