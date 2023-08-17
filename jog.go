package jog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type JLog struct {
	ZapLogger *zap.Logger
	ZapSugar  *zap.SugaredLogger
}

type LogLevel int

const (
	LevelProd LogLevel = iota
	LevelDev
)

var instance *JLog

func New(level ...LogLevel) *JLog {
	ll := LevelProd
	if len(level) == 0 {
		ll = LevelDev
	}
	jl := &JLog{}
	zapCfg := zap.NewProductionConfig()
	if ll == LevelDev {
		zapCfg = zap.NewDevelopmentConfig()
		zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	jl.ZapLogger, _ = zapCfg.Build()
	jl.ZapSugar = jl.ZapLogger.Sugar()
	instance = jl
	return jl
}

func GetLogger() *zap.Logger {
	if instance == nil {
		New()
	}
	return instance.ZapLogger
}

func GetSugar() *zap.SugaredLogger {
	if instance == nil {
		New()
	}
	return instance.ZapSugar
}

func (jl *JLog) Close() {
	jl.ZapLogger.Sync()
}

func Close() {
	if instance == nil {
		return
	}
	instance.Close()
}

func Debug(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Debugf(template, args...)
}

func Info(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Info(args...)
}

func Infof(template string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Infof(template, args...)
}

func Warn(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Warnf(template, args...)
}

func Error(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Fatalf(template, args...)
}

func Panic(args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	if instance == nil {
		New()
	}
	instance.ZapSugar.Panicf(template, args...)
}
