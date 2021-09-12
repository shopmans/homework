package main

import (
	// 标准包
	"database/sql"

	// 第三方包

	// 内部包
	"gotraining3/internal/week1"
	"gotraining3/internal/week2"
	"gotraining3/internal/week4"
	"gotraining3/internal/week5"
	week6 "gotraining3/internal/week6/pkg"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go3training")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 第一周作业
	// 答：
	// 当dao遇到sql.ErrNoRows时不需要wrap这个错误抛给上层
	// 1，sql.ErrNoRows不是一个错误
	// 2，dao查询结果返回sql.ErrNoRows本身也是一个查询结果
	// 3，上层业务是否wrap sql.ErrNoRows与具体业务相关，比如业务是查询一个user当收到sql.ErrNoRows本身也
	//    是一个正确结果不是错误无须wrap, 当修改用户时收到sql.ErrNoRows时则是一个错误因为修改用户业务
	//    无法执行下去这时需要wrap这个错误记录为什么修改用户失败了
	week1.Start(db)

	// 第二周作业
	week2.Start()

	// 第三周作业
	// pictureServer := wire.InitPictureService()
	// grpcServer := grpc.NewServer(grpc.Address(":9000"))
	// api.RegisterHomeworkWeek3Server(grpcServer, pictureServer)
	// app := kratos.New(
	// 	kratos.Name("homework_week3"),
	// 	kratos.Version("v1.1"),
	// 	kratos.Server(
	// 		grpcServer,
	// 	),
	// )

	// if err := app.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	// 第四周作业
	week4.Start()

	// 第五周作业
	week5.Start()

	// 第六周作业
	week6.Start()
}
