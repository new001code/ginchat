package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Log struct{}

func (l *Log) Debugln(s ...any) {
	fmt.Fprintln(gin.DefaultWriter, s...)
}

func (l *Log) Errorln(s ...any) {
	fmt.Fprintln(gin.DefaultErrorWriter, s...)
}

func (l *Log) Debugf(format string, s ...any) {
	fmt.Fprintf(gin.DefaultWriter, format, s...)
}

func (l *Log) Errorf(format string, s ...any) {
	fmt.Fprintf(gin.DefaultErrorWriter, format, s...)
}
