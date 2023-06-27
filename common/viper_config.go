package common

import (
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	logInit()
	confInit()
}

func logInit() {
	write := zapcore.AddSync(os.Stdout)
	//日志格式配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000000")
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, write, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar()
}

func confInit() {
	Logger.Info("start load configuration file")
	viper.SetConfigFile("./config/application.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Fatalf("fatal error reading configuration file: %s", err)
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		Logger.Warnf("config file changed: %s \n", e.Name)
	})
	Logger.Info("success load configuration file")
}
