package week2

import (
	// 标准包
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	// 第三方包
	"golang.org/x/sync/errgroup"
	// 内部包
)

func Start() {
	g, ctx := errgroup.WithContext(context.Background())
	serverCtx, cancel := context.WithCancel(ctx)
	httpserver1 := http.Server{Addr: "127.0.0.1:8001"}
	httpserver2 := http.Server{Addr: "127.0.0.1:8002"}
	httpserver3 := http.Server{Addr: "127.0.0.1:8003"}

	// 启动服务
	startServer1(serverCtx, g, &httpserver1)
	startServer1(serverCtx, g, &httpserver2)
	startServer1(serverCtx, g, &httpserver3)

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

func startServer1(ctx context.Context, g *errgroup.Group, server *http.Server) {
	g.Go(func() error {
		done := make(chan error)
		defer close(done)

		go func() {
			fmt.Println("server1 listen at 8001")
			done <- server.ListenAndServe()
		}()

		select {
		case err := <-done:
			// 返回一个非空err保证一个服务结束所有服务都结束
			if nil == err {
				err = fmt.Errorf("server1 exit")
			}
			return err
		case <-ctx.Done():
			return server.Shutdown(ctx)
		}
	})
}

func startServer2(ctx context.Context, g *errgroup.Group, server *http.Server) {
	g.Go(func() error {
		done := make(chan error)
		defer close(done)

		go func() {
			fmt.Println("server2 listen at 8002")
			done <- server.ListenAndServe()
		}()

		select {
		case err := <-done:
			// 返回一个非空err保证一个服务结束所有服务都结束
			if nil == err {
				err = fmt.Errorf("server2 exit")
			}
			return err
		case <-ctx.Done():
			return server.Shutdown(ctx)
		}
	})
}

func startServer3(ctx context.Context, g *errgroup.Group, server *http.Server) {
	g.Go(func() error {
		done := make(chan error)
		defer close(done)

		go func() {
			fmt.Println("server3 listen at 8003")
			done <- server.ListenAndServe()
		}()

		select {
		case err := <-done:
			// 返回一个非空err保证一个服务结束所有服务都结束
			if nil == err {
				err = fmt.Errorf("server3 exit")
			}
			return err
		case <-ctx.Done():
			return server.Shutdown(ctx)
		}
	})
}
