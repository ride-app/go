package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func New(dev bool, debug bool) *ZapLogger {
	encoderConfig := zap.NewProductionEncoderConfig()

	if dev {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.LevelKey = "severity"
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	if dev {
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       dev,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	if dev {
		zapConfig.Encoding = "console"
	}

	if debug {
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger := zap.Must(zapConfig.Build(zap.AddCallerSkip(1))).Sugar()

	return &ZapLogger{
		logger: logger,
	}
}

// func (l *LogrusLogger) Trace(args ...interface{}) {
// 	l.logger.Trace(args...)
// }

// func (l *LogrusLogger) Tracef(format string, args ...interface{}) {
// 	l.logger.Tracef(format, args...)
// }

func (l *ZapLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *ZapLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *ZapLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *ZapLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *ZapLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *ZapLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *ZapLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *ZapLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *ZapLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *ZapLogger) WithField(key string, value interface{}) Logger {
	return &ZapLogger{
		logger: l.logger.With(key, value),
	}
}

func (l *ZapLogger) WithFields(fields map[string]string) Logger {
	logger := l.logger
	for key, value := range fields {
		logger = logger.With(key, value)
	}
	return &ZapLogger{
		logger: logger,
	}
}

func (l *ZapLogger) WithError(err error) Logger {
	return &ZapLogger{
		logger: l.logger.With(zap.Error(err)),
	}
}

func (l *ZapLogger) WithContext(ctx context.Context) Logger {
	return l
}
