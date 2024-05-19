package jog

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestLogger(t *testing.T) {
	// Creating a zap logger for testing purposes
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	logger, err := config.Build()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()
	jl := &jLog{
		zapLogger: logger,
		zapSugar:  sugar,
	}

	// Test cases
	t.Run("Debug", func(t *testing.T) {
		jl.Debug("Debug message")
	})

	t.Run("Log", func(t *testing.T) {
		jl.Log("Log message")
	})

	t.Run("Info", func(t *testing.T) {
		jl.Info("Info message")
	})

	t.Run("Warn", func(t *testing.T) {
		jl.Warn("Warn message")
	})

	t.Run("Error", func(t *testing.T) {
		jl.Error("Error message")
	})

	t.Run("Fatal", func(t *testing.T) {
		// Note: Fatal will call os.Exit(1), which will terminate the test, so it's commented out
		// jl.Fatal("Fatal message")
	})

	t.Run("Panic", func(t *testing.T) {
		// Note: Panic will cause a panic, so it's commented out
		// jl.Panic("Panic message")
	})

	t.Run("Debugf", func(t *testing.T) {
		jl.Debugf("Debugf message: %s", "formatted")
	})

	t.Run("Logf", func(t *testing.T) {
		jl.Logf("Logf message: %s", "formatted")
	})

	t.Run("Infof", func(t *testing.T) {
		jl.Infof("Infof message: %s", "formatted")
	})

	t.Run("Warnf", func(t *testing.T) {
		jl.Warnf("Warnf message: %s", "formatted")
	})

	t.Run("Errorf", func(t *testing.T) {
		jl.Errorf("Errorf message: %s", "formatted")
	})

	t.Run("Fatalf", func(t *testing.T) {
		// Note: Fatalf will call os.Exit(1), which will terminate the test, so it's commented out
		// jl.Fatalf("Fatalf message: %s", "formatted")
	})

	t.Run("Panicf", func(t *testing.T) {
		// Note: Panicf will cause a panic, so it's commented out
		// jl.Panicf("Panicf message: %s", "formatted")
	})

	t.Run("Debugw", func(t *testing.T) {
		jl.Debugw("Debugw message", "key1", "value1")
	})

	t.Run("Infow", func(t *testing.T) {
		jl.Infow("Infow message", "key1", "value1")
	})

	t.Run("Warnw", func(t *testing.T) {
		jl.Warnw("Warnw message", "key1", "value1")
	})

	t.Run("Errorw", func(t *testing.T) {
		jl.Errorw("Errorw message", "key1", "value1")
	})

	t.Run("Fatalw", func(t *testing.T) {
		// Note: Fatalw will call os.Exit(1), which will terminate the test, so it's commented out
		// jl.Fatalw("Fatalw message", "key1", "value1")
	})

	t.Run("Panicw", func(t *testing.T) {
		// Note: Panicw will cause a panic, so it's commented out
		// jl.Panicw("Panicw message", "key1", "value1")
	})
}

func TestNew(t *testing.T) {
	logger := New(LevelDev)
	if logger == nil {
		t.Error("Expected logger instance, got nil")
	}
}

func TestGetLogger(t *testing.T) {
	New(LevelDev)
	logger := GetLogger()
	if logger == nil {
		t.Error("Expected logger instance, got nil")
	}
}

func TestGetZapSugaredLogger(t *testing.T) {
	New(LevelDev)
	sugar := GetZapSugaredLogger()
	if sugar == nil {
		t.Error("Expected Zap SugaredLogger instance, got nil")
	}
}

func TestGetZapLogger(t *testing.T) {
	New(LevelDev)
	zapLogger := GetZapLogger()
	if zapLogger == nil {
		t.Error("Expected Zap Logger instance, got nil")
	}
}

func TestClose(t *testing.T) {
	logger := New(LevelDev)
	logger.Close()
	// Closing should not raise an error
}

func TestLoggingFunctions(t *testing.T) {
	New(LevelDev)

	t.Run("Debug", func(t *testing.T) {
		Debug("Debug message")
	})

	t.Run("Log", func(t *testing.T) {
		Log("Log message")
	})

	t.Run("Info", func(t *testing.T) {
		Info("Info message")
	})

	t.Run("Warn", func(t *testing.T) {
		Warn("Warn message")
	})

	t.Run("Error", func(t *testing.T) {
		Error("Error message")
	})

	// Fatal and Panic tests are commented out to avoid terminating the test run

	t.Run("Debugf", func(t *testing.T) {
		Debugf("Debugf message: %s", "formatted")
	})

	t.Run("Logf", func(t *testing.T) {
		Logf("Logf message: %s", "formatted")
	})

	t.Run("Infof", func(t *testing.T) {
		Infof("Infof message: %s", "formatted")
	})

	t.Run("Warnf", func(t *testing.T) {
		Warnf("Warnf message: %s", "formatted")
	})

	t.Run("Errorf", func(t *testing.T) {
		Errorf("Errorf message: %s", "formatted")
	})

	t.Run("Debugw", func(t *testing.T) {
		Debugw("Debugw message", "key1", "value1")
	})

	t.Run("Infow", func(t *testing.T) {
		Infow("Infow message", "key1", "value1")
	})

	t.Run("Warnw", func(t *testing.T) {
		Warnw("Warnw message", "key1", "value1")
	})

	t.Run("Errorw", func(t *testing.T) {
		Errorw("Errorw message", "key1", "value1")
	})
}
