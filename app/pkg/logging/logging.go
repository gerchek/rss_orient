package logging

import (
	"fmt"
	"io"
//	"io/ioutil"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

const logFile = "logs/project.log"

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

	log.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := f.File
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

//	log.Out = ioutil.Discard

	f, err := os.OpenFile("logs/project.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// wrt := io.MultiWriter(os.Stdout, f)

	// log.SetOutput(wrt)

	// log.SetOutput(io.MultiWriter(f, os.Stdout))

//	log.SetOutput(io.Discard)

	log.AddHook(&writerHook{
		Writer:    []io.Writer{f},
		LogLevels: logrus.AllLevels,
	})

}

func Log() *logrus.Logger {
	return log
}
