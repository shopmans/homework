package pkg

import (
	// 标准包
	"context"
	"fmt"
	"os"
	"os/signal"

	// 第三方包
	"golang.org/x/sync/errgroup"
	// 内部包
	"gotraining3/internal/week6/service"
)

func Start() {
	g, ctx := errgroup.WithContext(context.Background())
	serverCtx, cancel := context.WithCancel(ctx)
	fixLengthServer := service.NewFixLengthDecode(serverCtx, 4, "8881")
	delimiterBasedServer := service.NewDeLimiterBasedDecode(serverCtx, "$#", "8882")
	LengthFieldBasedServer := service.NewLengthFieldBasedFrameDecoder(serverCtx, 24, 0, 2, 0, true, "8883")

	// 启动服务
	g.Go(func() error {
		return fixLengthServer.Start()
	})
	g.Go(func() error {
		return delimiterBasedServer.Start()
	})
	g.Go(func() error {
		return LengthFieldBasedServer.Start()
	})

	// 监听信号
	go func() {
		done := make(chan os.Signal)
		defer close(done)
		signal.Notify(done, os.Interrupt, os.Kill)
		<-done
		cancel()
	}()

	err := g.Wait()
	fmt.Println(err)
	fmt.Println(serverCtx.Err())
}
