package main

import (
	"log"
	"net/http"
	"time"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/yakushou730/blog-service/pkg/logger"

	"github.com/yakushou730/blog-service/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/yakushou730/blog-service/global"
	"github.com/yakushou730/blog-service/internal/routers"
	"github.com/yakushou730/blog-service/pkg/setting"
)

func init() {
	err := SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

}

// @title 部落格系統
// @version 1.0
// @description Go 語言程式設計之旅：一起用 Go 做專案
// @termsOfService https://gihub.com/go-programming-tour-book
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:              ":" + global.ServerSetting.HttpPort,
		Handler:           router,
		ReadTimeout:       global.ServerSetting.ReadTimeout,
		ReadHeaderTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	global.Logger.Infof("%s: go-programming-tour-book/%s", "yakushou730", "blog-service")

	s.ListenAndServe()
}

func SetupSetting() error {
	systemSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = systemSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = systemSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = systemSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
