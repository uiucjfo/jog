package jog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel int

const (
	LevelProd LogLevel = iota
	LevelDev
)

var instance Logger

func New(level ...LogLevel) Logger {
	ll := LevelProd
	if len(level) == 0 {
		ll = LevelDev
	} else {
		ll = level[0]
	}
	j := &jLog{}
	zapCfg := zap.NewProductionConfig()
	if ll == LevelDev {
		zapCfg = zap.NewDevelopmentConfig()
		zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	j.zapLogger, _ = zapCfg.Build()
	j.zapSugar = j.zapLogger.Sugar().WithOptions(zap.AddCallerSkip(2))
	instance = j
	return j
}

func GetLogger() Logger {
	return instance
}

func GetZapSugaredLogger() *zap.SugaredLogger {
	if instance == nil {
		New()
	}
	return instance.GetZapSugarLogger()
}

func GetZapLogger() *zap.Logger {
	if instance == nil {
		New()
	}
	return instance.GetZapLogger()
}

func Close() {
	instance.Close()
}

func Debug(args ...any) {
	if instance == nil {
		New()
	}
	instance.Debug(args...)
}

func Log(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Log(args...)
}

func Info(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Info(args...)
}

func Warn(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Warn(args...)
}

func Error(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Error(args...)
}

func Fatal(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Fatal(args...)
}

func Panic(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Panic(args...)
}

func Debugf(format string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Debugf(format, args...)
}

func Logf(format string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Logf(format, args...)
}

func Infof(format string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Panicf(format, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Errorw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Fatalw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	if instance == nil {
		New()
	}
	instance.Panicw(msg, keysAndValues...)
}
