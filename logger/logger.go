package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

// Info alerts valuable info checkpoints
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Fatal alerts breaking changes that will terminate the app
func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}

// Debug alerts useful buging messages that can happen
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error alerts when the app is not used as designed causing an error
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
