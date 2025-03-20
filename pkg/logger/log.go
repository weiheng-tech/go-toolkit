package logger

import (
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/weiheng-tech/go-toolkit/pkg/utils"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func ensureLogPath(path string) {
	if existed, _ := utils.PathExists(filepath.Dir(path)); existed {
		return
	} else {
		_ = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	}
}

func newLfsHook(config *LogOptions, format log.Formatter) (log.Hook, io.Writer, error) {
	baseLogPath := config.Path
	ensureLogPath(baseLogPath)

	writer := &lumberjack.Logger{
		Filename:   baseLogPath,
		MaxSize:    config.RotationSize,
		MaxBackups: config.RotationCount,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	lfsHook := lfshook.NewHook(writer, format)

	return lfsHook, writer, nil
}

func InitLog(config *LogOptions) {
	log.SetReportCaller(true)
	log.AddHook(&CallerHook{})

	level := config.Level
	parsedLevel, err := log.ParseLevel(level)
	if err != nil {
		log.Warning("Failed to parse log level, setting to default (debug)")
		log.SetLevel(log.DebugLevel)
		return
	}

	log.SetLevel(parsedLevel)
	log.SetOutput(os.Stdout)

	var formatter log.Formatter
	if config.Format == "json" {
		formatter = &log.JSONFormatter{}
		log.SetFormatter(formatter)
	} else {
		formatter = &log.TextFormatter{
			FullTimestamp: true,
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				return "", ""
			},
			DisableQuote: config.DisableQuote,
		}
		log.SetFormatter(formatter)
	}

	if len(config.Path) > 0 {
		fileHook, _, err := newLfsHook(config, formatter)
		if err != nil {
			log.Errorf("config local file system for logger error: %v", err)
		} else {
			log.AddHook(fileHook)
		}
	}

}
