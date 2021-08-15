package week4

import (
	// 标准包
	"fmt"
	"strconv"
	"time"

	// 第三方包
	// 内部包
	roll "gotraining3/internal/week4/pkg"
)

func Start() {
	// 初始化一个滑动窗口，10个bucket 第一个bucket 点 100 毫秒
	bucketState := roll.NewBucketState(10, 100)

	for i := 0; i < 100; i++ {
		time.Sleep(50)
		err, bucket := bucketState.GetCurrentBucket()
		if nil != err {
			fmt.Println(err.Error())
			return
		}
		bucket.Add()
	}

	fmt.Println("统计滑动窗口计数总数： " + strconv.Itoa(bucketState.GetWindowsCount()))
}
