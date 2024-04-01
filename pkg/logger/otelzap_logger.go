package logger

import (
	"context"
	"os"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type OtelZapLogger struct {
	logger        *otelzap.SugaredLogger
	loggerWithCtx *otelzap.SugaredLoggerWithCtx
}

func NewOtel(dev bool, debug bool) *OtelZapLogger {
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

	zapLogger := zap.Must(zapConfig.Build(zap.AddCallerSkip(1)))

	logger := otelzap.New(zapLogger).Sugar()

	return &OtelZapLogger{
		logger: logger,
	}
}

// func (l *LogrusLogger) Trace(args ...interface{}) {
// 	l.logger.Trace(args...)
// }

// func (l *LogrusLogger) Tracef(format string, args ...interface{}) {
// 	l.logger.Tracef(format, args...)
// }

func (l *OtelZapLogger) Debug(args ...interface{}) {
	if l.loggerWithCtx != nil {
		l.loggerWithCtx.Debug(args...)
	} else {
		l.logger.Debug(args...)
	}
}

func (l *OtelZapLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *OtelZapLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *OtelZapLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *OtelZapLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *OtelZapLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *OtelZapLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *OtelZapLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *OtelZapLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *OtelZapLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *OtelZapLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *OtelZapLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *OtelZapLogger) WithField(key string, value interface{}) Logger {
	return &OtelZapLogger{
		logger: l.logger.With(key, value),
	}
}

func (l *OtelZapLogger) WithFields(fields map[string]string) Logger {
	logger := l.logger
	for key, value := range fields {
		logger = logger.With(key, value)
	}
	return &OtelZapLogger{
		logger: logger,
	}
}

func (l *OtelZapLogger) WithError(err error) Logger {
	return &OtelZapLogger{
		logger: l.logger.With(zap.Error(err)),
	}
}

func (l *OtelZapLogger) WithContext(ctx context.Context) Logger {
	lctx := l.logger.Ctx(ctx)
	return &OtelZapLogger{
		logger:        l.logger,
		loggerWithCtx: &lctx,
	}
}
