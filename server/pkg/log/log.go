package log

import (
	"log/syslog"

	"github.com/tchap/zapext/zapsyslog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewToSyslog returns new syslogger for Zap
func NewToSyslog(tag, version string, prod bool) *zap.SugaredLogger {
	writer, err := syslog.New(syslog.LOG_ERR|syslog.LOG_LOCAL0, tag)
	if err != nil {
		panic("failed to set up syslog: " + err.Error())
	}
	encoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapsyslog.NewCore(zapcore.DebugLevel, encoder, writer)
	l := zap.New(core, zap.Development(), zap.AddStacktrace(zapcore.ErrorLevel))
	return l.Sugar().With("version", version).With("tag", tag)
}

// NewToFile creates a new file logger
func NewToFile(filename, tag, version string, prod bool) *zap.SugaredLogger {
	var conf zap.Config
	if !prod {
		conf = zap.NewDevelopmentConfig()
		conf.OutputPaths = []string{
			filename,
		}
	} else {
		conf = zap.NewProductionConfig()
		conf.OutputPaths = []string{
			filename,
		}
	}

	l, err := conf.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.FatalLevel),
	)
	if err != nil {
		panic(err)
	}
	return l.Sugar().With("version", version).With("tag", tag)
}

// NewToStdOut creates a new file logger
func NewToStdOut(tag, version string, prod bool) *zap.SugaredLogger {

	if prod {
		config := zap.NewProductionConfig()
		l, err := config.Build()
		if err != nil {
			panic("can't initialize zap logger: " + err.Error())
		}
		return l.Sugar().With("version", version).With("tag", tag)
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, err := config.Build()
	if err != nil {
		panic("can't initialize zap logger: " + err.Error())
	}
	return l.Sugar().With("version", version).With("tag", tag)
}
