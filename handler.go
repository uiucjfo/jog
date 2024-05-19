package jog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type jLog struct {
	zapLogger *zap.Logger
	zapSugar  *zap.SugaredLogger
}

func (jl *jLog) Debug(args ...any) {
	jl.zapSugar.Debug(args...)
}

func (jl *jLog) Log(args ...any) {
	jl.zapSugar.Log(zapcore.InfoLevel, args...)
}

func (jl *jLog) Info(args ...any) {
	jl.zapSugar.Info(args...)
}

func (jl *jLog) Warn(args ...any) {
	jl.zapSugar.Warn(args...)
}

func (jl *jLog) Error(args ...any) {
	jl.zapSugar.Error(args...)
}

func (jl *jLog) Fatal(args ...any) {
	jl.zapSugar.Fatal(args...)
}

func (jl *jLog) Panic(args ...any) {
	jl.zapSugar.Panic(args...)
}

func (jl *jLog) Debugf(format string, args ...any) {
	jl.zapSugar.Debugf(format, args...)
}

func (jl *jLog) Logf(format string, args ...any) {
	jl.zapSugar.Logf(zapcore.InfoLevel, format, args...)
}

func (jl *jLog) Infof(format string, args ...any) {
	jl.zapSugar.Infof(format, args...)
}

func (jl *jLog) Warnf(format string, args ...any) {
	jl.zapSugar.Warnf(format, args...)
}

func (jl *jLog) Errorf(format string, args ...any) {
	jl.zapSugar.Errorf(format, args...)
}

func (jl *jLog) Fatalf(format string, args ...any) {
	jl.zapSugar.Fatalf(format, args...)
}

func (jl *jLog) Panicf(format string, args ...any) {
	jl.zapSugar.Panicf(format, args...)
}

func (jl *jLog) Debugw(msg string, keysAndValues ...any) {
	jl.zapSugar.Debugw(msg, keysAndValues...)
}

func (jl *jLog) Infow(msg string, keysAndValues ...any) {
	jl.zapSugar.Infow(msg, keysAndValues...)
}

func (jl *jLog) Warnw(msg string, keysAndValues ...any) {
	jl.zapSugar.Warnw(msg, keysAndValues...)
}

func (jl *jLog) Errorw(msg string, keysAndValues ...any) {
	jl.zapSugar.Errorw(msg, keysAndValues...)
}

func (jl *jLog) Fatalw(msg string, keysAndValues ...any) {
	jl.zapSugar.Fatalw(msg, keysAndValues...)
}

func (jl *jLog) Panicw(msg string, keysAndValues ...any) {
	jl.zapSugar.Panicw(msg, keysAndValues...)
}

func (jl *jLog) GetZapLogger() *zap.Logger {
	return jl.zapLogger
}

func (jl *jLog) GetZapSugarLogger() *zap.SugaredLogger {
	return jl.zapSugar
}

func (jl *jLog) Close() {
	jl.zapLogger.Sync()
}
