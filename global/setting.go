package global

import (
	"github.com/go-grogramming-tour-book/blog-service/pkg/logger"
	"github.com/go-grogramming-tour-book/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS

	Logger *logger.Logger
)
