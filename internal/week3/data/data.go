package data

import (
	// 标准包
	"context"
	// 第三方包
	"go.mongodb.org/mongo-driver/mongo"
	// 内部包
	"gotraining3/internal/week3/biz"
)

type pictrueRepo struct {
	db *mongo.Database
}

func NewPictureRepo(db *mongo.Database) biz.PictureRepo {
	return pictrueRepo{db: db}
}

func (pr pictrueRepo) Modify(p *biz.Picture) error {
	_, err := pr.db.Collection("picture").InsertOne(context.TODO(), p)
	return err
}
