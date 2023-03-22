package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"assessment/internal/core/config"
)

type LoggerInt interface {
	Initialize()
	Debug(args ...interface{})
	Debugf(msg string, args ...interface{})
	Debugw(msg string, args ...interface{})
	Info(args ...interface{})
	Infof(msg string, args ...interface{})
	Infow(msg string, args ...interface{})
	Warn(args ...interface{})
	Warnf(msg string, args ...interface{})
	Warnw(msg string, args ...interface{})
	Error(args ...interface{})
	Errorf(msg string, args ...interface{})
	Errorw(msg string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(msg string, args ...interface{})
	Fatalw(msg string, args ...interface{})
}

type Logger struct {
	appLogger *zap.SugaredLogger
	cfg       *config.Config
}

func NewLogger(cfg *config.Config) *Logger {
	return &Logger{
		cfg: cfg,
	}
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func (l *Logger) Initialize() {
	logLevel, ok := loggerLevelMap[l.cfg.Logger.Level]
	if !ok {
		logLevel = zapcore.DebugLevel
	}

	atom := zap.NewAtomicLevelAt(logLevel)
	encoderCfg := zap.NewProductionEncoderConfig()
	zapOutput := zapcore.Lock(os.Stdout)

	jsonEncoder := zapcore.NewJSONEncoder(encoderCfg)
	zapCore := zapcore.NewCore(jsonEncoder, zapOutput, atom)

	l.appLogger = zap.New(zapCore).Sugar()
}

func (l *Logger) Debug(args ...interface{}) {
	l.appLogger.Debug(args)
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.appLogger.Debugf(msg, args...)
}

func (l *Logger) Debugw(msg string, args ...interface{}) {
	l.appLogger.Debugw(msg, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.appLogger.Info(args...)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.appLogger.Infof(msg, args...)
}

func (l *Logger) Infow(msg string, args ...interface{}) {
	l.appLogger.Infow(msg, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.appLogger.Warn(args...)
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.appLogger.Warnf(msg, args...)
}

func (l *Logger) Warnw(msg string, args ...interface{}) {
	l.appLogger.Warnw(msg, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.appLogger.Error(args...)
}

func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.appLogger.Errorf(msg, args...)
}

func (l *Logger) Errorw(msg string, args ...interface{}) {
	l.appLogger.Errorw(msg, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.appLogger.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.appLogger.Fatalf(template, args...)
}

func (l *Logger) Fatalw(template string, args ...interface{}) {
	l.appLogger.Fatalw(template, args...)
}
