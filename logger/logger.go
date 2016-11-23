package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/tonyalaribe/monitor-server/config"
	"github.com/tonyalaribe/monitor-server/constants"
)

var log = logrus.New()

func Init() {

	file, err := os.OpenFile(config.Get().LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	log.Out = file

	log.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: true}
}

func getDebugInfo(callerDepth int) string {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[callerDepth])
	fs := strings.Split(f.Name(), "/")
	fn := fs[len(fs)-1]
	file, line := f.FileLine(pc[callerDepth] - 1)
	files := strings.Split(file, "/")
	file = files[len(files)-1]
	return fmt.Sprintf("%s(%s:%d)", fn, file, line)
}

func with(c *gin.Context) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		constants.LOCATION: getDebugInfo(2),
		//constants.PATH:       c.Request.URL.Path,
		//constants.METHOD:     c.Request.Method,
	})
}

func Info(c *gin.Context, args ...interface{}) {
	with(c).Infoln(args)
}

func Debug(c *gin.Context, args ...interface{}) {
	with(c).Debugln(args)
}

func Warn(c *gin.Context, args ...interface{}) {
	with(c).Warnln(args)
}

func Error(c *gin.Context, args ...interface{}) {
	with(c).Errorln(args)
}

// from other places than resources
func ErrorWithoutContext(args ...interface{}) {
	log.WithField(constants.LOCATION, getDebugInfo(1)).Errorln(args)
}
