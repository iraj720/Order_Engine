package confs

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger() *zap.Logger {

	logLevel := os.Getenv("LOG_LEVEL")
	devMod := os.Getenv("DEVELOPMENT_MOD")

	zapLevel := zap.DebugLevel
	switch strings.ToLower(logLevel) {
	case "info":
		zapLevel = zap.InfoLevel
	case "warn":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	case "fatal":
		zapLevel = zap.FatalLevel
	case "debug":
		zapLevel = zap.DebugLevel

	}
	development := false
	if strings.EqualFold(devMod, "true") {
		development = true
	}
	conf := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapLevel),
		Development:      development,
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			CallerKey:    "caller",
			FunctionKey:  "func",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var err error
	l, err := conf.Build()
	if err != nil {
		panic("failed to initiate logger")
	}

	return l
}
