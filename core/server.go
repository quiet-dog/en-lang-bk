package core

import (
	"context"
	"en-lang-bk/global"
	"en-lang-bk/initialize"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
)

func RunServer() {

	global.DB = initialize.Gorm()

	addr := fmt.Sprintf("%s:%d", global.CONFIG.System.Host, global.CONFIG.System.Port)
	router := initialize.InitRouters()
	srv := http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.LOG.Error("listen", zap.Error(err))
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.LOG.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.LOG.Error("Server Shutdown", zap.Error(err))
	}
	global.LOG.Info("Server exiting")
}
