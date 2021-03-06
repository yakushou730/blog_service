package global

import (
	"github.com/yakushou730/blog-service/pkg/logger"
	"github.com/yakushou730/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	JWTSetting      *setting.JWTSettings
	Logger          *logger.Logger
)
