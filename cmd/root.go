package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hangter-lt/easy_mock/global"
	"github.com/hangter-lt/easy_mock/router"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "easy_mock",
		Short: "api mock",
		Run: func(cmd *cobra.Command, args []string) {
			defer global.DB.Close()

			// 启动管理端
			srv1 := &http.Server{
				Addr:    "0.0.0.0:7001",
				Handler: router.RouterManage(),
			}

			go func() {
				if err := srv1.ListenAndServe(); err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						log.Fatal(errors.New("user close server"))
					} else {
						fmt.Printf("err: %v\n", err)
						log.Fatal("server exited")
					}
				}
			}()

			// 启动接收端
			srv2 := &http.Server{
				Addr:    "0.0.0.0:7002",
				Handler: router.RouterCall(),
			}
			go func() {
				if err := srv2.ListenAndServe(); err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						log.Fatal(errors.New("user close server"))
					} else {
						log.Fatal("server exited")
					}
				}
			}()

			quit := make(chan os.Signal, 2)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := srv1.Shutdown(ctx); err != nil {
				log.Fatal("server exited")
			}

			if err := srv2.Shutdown(ctx); err != nil {
				log.Fatal("server exited")
			}
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
