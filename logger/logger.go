package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	logConfig := zap.Config{
		//no file in std
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		//to elastic search
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:   "level",
			TimeKey:    "time",
			MessageKey: "msg",
			//2006-01-02T15:04:05.000Z07002
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	//init login system result
	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
	//Log.Info("algoo")
}

//varying fn
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	log.Sync()
}

func GetLogger() *zap.Logger {
	return log
}
