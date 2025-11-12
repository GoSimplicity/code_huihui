package di

import (
	"os"
	"path/filepath"
	"time"

	"github.com/samber/do"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(i *do.Injector) (*zap.Logger, error) {
	v := do.MustInvoke[*viper.Viper](i)
	logDir := v.GetString("log.dir")
	currentTime := time.Now().Format("2006-01-02")
	logFile := filepath.Join(logDir, "code-huihui-"+currentTime+".log")

	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic("无法创建日志目录")
	}

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
		LocalTime:  true,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	consoleWriter := zapcore.AddSync(os.Stdout)
	logLevel := getLogLevel(v.GetString("log.level"))

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleWriter, logLevel),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileWriter, logLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger, nil
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel // 默认使用 Info 级别
	}
}
