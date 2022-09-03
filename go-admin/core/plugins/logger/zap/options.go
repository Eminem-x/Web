package zap

import (
	"go-admin/core/logger"
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	logger.Options
}

type callerSkipKey struct{}

func WithCallerSkip(i int) logger.Option {
	return logger.SetOption(callerSkipKey{}, i)
}

type configKey struct{}

// WithConfig pass zap.config to logger
func WithConfig(c zap.Config) logger.Option {
	return logger.SetOption(configKey{}, c)
}

type encoderConfigKey struct{}

// WithEncoderConfig pass zapcore.EncodeConfig to logger
func WithEncoderConfig(c zapcore.EncoderConfig) logger.Option {
	return logger.SetOption(encoderConfigKey{}, c)
}

type namespaceKey struct{}

func WithNameSpace(namespace string) logger.Option {
	return logger.SetOption(namespaceKey{}, namespace)
}

type writerKey struct{}

func WithOutput(out io.Writer) logger.Option {
	return logger.SetOption(writerKey{}, out)
}
