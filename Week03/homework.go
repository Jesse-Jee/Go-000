package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

func ServerDemo1(ctx context.Context) error {
	srv := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil,
	}

	go func() {
		<-ctx.Done()
		log.Println("demo1 graceful shutdown")
		srv.Shutdown(context.Background())
	}()

	return srv.ListenAndServe()
}

func ServerDemo2(ctx context.Context) error {
	srv := http.Server{
		Addr:    "0.0.0.0:9090",
		Handler: nil,
	}

	go func() {
		<-ctx.Done()
		log.Println("demo2 graceful shutdown")
		srv.Shutdown(context.Background())
	}()
	return srv.ListenAndServe()
}

func listenSig(ctx context.Context) error {
	sig := make(chan os.Signal)
	signal.Notify(sig)
	for {
		select {
		case stop := <-sig:
			return fmt.Errorf("recieved stop signal %v", stop)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return ServerDemo1(ctx)
	})

	group.Go(func() error {
		return ServerDemo2(ctx)
	})

	group.Go(func() error {
		return listenSig(ctx)
	})

	if err := group.Wait(); err != nil {
		fmt.Println("errGroup get error ", err)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("see you...")
}
