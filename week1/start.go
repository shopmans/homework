package week1

import (
	// 标准包
	"database/sql"
	"fmt"

	// 第三方包

	// 内部包
	"gotraining3/week1/business"
)

// 答：
// 当dao遇到sql.ErrNoRows时不需要wrap这个错误抛给上层
// 1，sql.ErrNoRows不是一个错误
// 2，dao查询结果返回sql.ErrNoRows本身也是一个查询结果
// 3，上层业务是否wrap sql.ErrNoRows与具体业务相关，比如业务是查询一个user当收到sql.ErrNoRows本身也
//    是一个正确结果不是错误无须wrap, 当修改用户时收到sql.ErrNoRows时则是一个错误因为修改用户业务
//    无法执行下去这时需要wrap这个错误记录为什么修改用户失败了

func Start(db *sql.DB) {
	// 查询
	user, err := business.Query(db, 1)
	if nil != err {
		if nil == user {
			fmt.Println("用户不存在")
		} else {
			fmt.Printf("查询到用户，%s/n/s", user.Name)
		}
	}
	// 修改
	user.Name = "凡尔赛"
	user, err = business.Modify(db, user)
	if nil != err {
		fmt.Printf("错误：%+v", err)
	}
}
