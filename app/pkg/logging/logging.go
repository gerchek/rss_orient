package logging

import (
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"runtime"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

var log *logrus.Logger

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}

	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func init() {
	log = logrus.New()
	// log to console and file
	log.SetReportCaller(true)

	log.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			// filename := f.File
			_, filename := path.Split(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	log.Out = ioutil.Discard

	//rl, err := rotatelogs.New("/home/ata/tps/rss_orient/app/cmd/project/logs/rss_log.%Y-%m-%d")
	 rl, err := rotatelogs.New("/var/www/rss/app/cmd/project/logs/rss_log.%Y-%m-%d")

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	//log.SetOutput(io.Discard)

	log.AddHook(&writerHook{
		Writer:    []io.Writer{rl},
		LogLevels: logrus.AllLevels,
	})

}

func Log() *logrus.Logger {
	return log
}
