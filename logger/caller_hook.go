package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type CallerHook struct {
}

func (h *CallerHook) Levels() []log.Level {
	return log.AllLevels
}

func (h *CallerHook) Fire(entry *log.Entry) error {

	if entry.Level <= log.InfoLevel {
		if entry.HasCaller() {
			if entry.Level == log.ErrorLevel {
				entry.Data["source"] = fmt.Sprintf("%s:%d:%s()", entry.Caller.File, entry.Caller.Line, entry.Caller.Function)
			}
		}
	}
	return nil
}
