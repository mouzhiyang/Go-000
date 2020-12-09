package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func HttpServer(ctx context.Context, addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	})
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		log.Println("server stop")
		server.Shutdown(ctx)
	}()
	return server.ListenAndServe()
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	//两个服务
	g.Go(func() error {
		if err := HttpServer(ctx, "8080"); err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		if err := HttpServer(ctx, "9090"); err != nil {
			return err
		}
		return nil
	})

	//监听信号
	g.Go(func() error {
		q := make(chan os.Signal)
		signal.Notify(q, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-q:
				return errors.New("exit")
			}
		}
	})

	if err := g.Wait(); err != nil {
		log.Println("退出")
	}
}
