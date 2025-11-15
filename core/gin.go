package core

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"run/global"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func RunServer() {
	global.Log.Info("starting server")

	var address string
	if global.Config.System.Port <= 0 {
		global.Log.Warn("server port is invalid", zap.Int("port", global.Config.System.Port))
		address = fmt.Sprintf(":%d", DefaultPort)
		global.Log.Warn("use default port", zap.Int("port", DefaultPort))
	} else {
		global.Log.Info("server port is valid", zap.Int("port", global.Config.System.Port))
		address = fmt.Sprintf(":%d", global.Config.System.Port)
	}

	svr := &http.Server{
		Addr:           address,
		Handler:        Routers(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := svr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal("server start error", zap.String("port", address), zap.Error(err))
			os.Exit(1)
		}
	}()

	zap.L().Info("server started", zap.String("address", "localhost"+address))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		zap.L().Fatal("server shutdown error", zap.Error(err))
	}
	zap.L().Info("server exited")
}
