package logcaller

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

//CallerHook it's for logrus
type CallerHook struct {
}

//Fire logurs hook interface
func (hook *CallerHook) Fire(entry *logrus.Entry) error {
	entry.Data["caller"] = hook.caller(5)
	return nil
}

//Levels logurs hook interface
func (hook *CallerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (hook *CallerHook) caller(skip int) string {
	if _, file, line, ok := runtime.Caller(skip); ok {
		return strings.Join([]string{filepath.Base(file), strconv.Itoa(line)}, ":")
	}
	return "?"
}
