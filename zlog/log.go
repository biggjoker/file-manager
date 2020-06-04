package zlog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	zap "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_zlog *zap.SugaredLogger
)

func InitLog(debug bool) io.Writer {
	path := getFilePath()
	writer, _ := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithLinkName(path),
	)

	w := zapcore.AddSync(writer)
	var (
		config zapcore.EncoderConfig
		core   zapcore.Core
	)
	if debug {
		config = zap.NewDevelopmentEncoderConfig()
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			w,
			zap.NewAtomicLevelAt(zap.DebugLevel),
		)
	} else {
		config = zap.NewProductionEncoderConfig()
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			w,
			zap.NewAtomicLevelAt(zap.InfoLevel),
		)
	}

	_logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	_zlog = _logger.Sugar()
	return w
}

func getFilePath() string {
	logfile := getCurrentDirectory() + "/logs/" + getAppname() + ".log"
	return logfile
}

func getAppname() string {
	full := os.Args[0]
	full = strings.Replace(full, "\\", "/", -1)
	splits := strings.Split(full, "/")
	if len(splits) >= 1 {
		name := splits[len(splits)-1]
		name = strings.TrimSuffix(name, ".exe")
		return name
	}

	return ""
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Debug(args ...interface{}) {
	_zlog.Debug(args...)
}

func Info(args ...interface{}) {
	_zlog.Info(args...)
}

func Warn(args ...interface{}) {
	_zlog.Warn(args...)
}

func Error(args ...interface{}) {
	_zlog.Error(args...)
}

func DPanic(args ...interface{}) {
	_zlog.DPanic(args...)
}

func Panic(args ...interface{}) {
	_zlog.Panic(args...)
}

func Fatal(args ...interface{}) {
	_zlog.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	_zlog.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	_zlog.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	_zlog.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	_zlog.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	_zlog.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	_zlog.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	_zlog.Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	_zlog.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	_zlog.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	_zlog.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	_zlog.Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	_zlog.DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	_zlog.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	_zlog.Fatalw(msg, keysAndValues...)
}
