package util

import (
	"log"

	"github.com/gin-gonic/gin"
)

type MyLog struct {
}

var (
	debugLogger *log.Logger
	errorLogger *log.Logger
)

func init() {
	debugLogger = log.New(gin.DefaultWriter, "[DEBUG]", log.Lshortfile|log.LstdFlags)
	errorLogger = log.New(gin.DefaultErrorWriter, "[ERROR]", log.Lshortfile|log.LstdFlags)
}

func (l *MyLog) Debugln(s ...any) {
	debugLogger.Println(s...)
}

func (l *MyLog) Errorln(s ...any) {
	errorLogger.Println(s...)
}

func (l *MyLog) Debugf(format string, s ...any) {
	debugLogger.Printf(format, s...)
}

func (l *MyLog) Errorf(format string, s ...any) {
	errorLogger.Printf(format, s...)
}
