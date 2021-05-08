package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})
	serv := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// 启动 http server
	g.Go(func() error {
		fmt.Println("start")
		go func() {
			<-ctx.Done()
			fmt.Println("ctx done")
			serv.Shutdown(ctx)
		}()
		return serv.ListenAndServe()
	})

	// 模拟服务错误退出
	g.Go(func() error {
		fmt.Println("10s后生成error")
		time.Sleep(10 * time.Second)
		return errors.New("error test。。。")
	})

	// signal信号
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-sigs:
			return fmt.Errorf("termin signal: %v", sigs)
		}
	})

	err := g.Wait()
	fmt.Println(err)
}
