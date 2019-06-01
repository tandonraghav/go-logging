package logging

import (
	"github.com/tandonraghav/go-logging/utilities"
	"github.com/tandonraghav/go-logging/web/config"
	"context"
	_  "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
	This class is used for Application wide Logging. All Logging statements should use below functions
	Logrus should not be used directly in any place.
*/
var log *logrus.Logger
var reqIDFn = utilities.GetRequestID

type contextLogger struct {
	entry *logrus.Entry
}

type (
	loggerKey struct{}
)


func InitializeLogger() {
	lvl := config.GetLogLevel()
	l, _ := logrus.ParseLevel(lvl)
	log = logrus.New()
	keys:=config.GetLogKeys()
	log.SetReportCaller(true)
	log.SetFormatter(&Formatter{
		HideKeys:      true,
		FieldsOrder:   keys,
		ShowFullLevel: true,
	})
	log.SetLevel(l)
	logFileName := config.GetLogFileName()
	log.SetOutput(&lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
}

func newContextLogger(ctx context.Context) *contextLogger {
	keys:=config.GetLogKeys()
	entry:=logrus.NewEntry(log)
	for k := range keys{
		entry=entry.WithField(keys[k],utilities.GetValueFromReqContext(ctx,keys[k]))
	}
	return &contextLogger{
		entry: entry,
	}
}

func GetLogger(ctx context.Context) *logrus.Entry {
	if ctx != nil {
		logger := ctx.Value(loggerKey{})
		if logger == nil {
			return logrus.NewEntry(log)
		}
		l, ok := logger.(*contextLogger)
		if ok {
			return l.entry
		}
	}
	return logrus.NewEntry(log)
}

func WithLogger(ctx context.Context) context.Context {
	return context.WithValue(ctx, loggerKey{}, newContextLogger(ctx))
}

func getContextValue(ctx context.Context, field interface{}) (result string) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("Exception while getting Ctx Value ", r)
			result = "NA"
		}

	}()
	if ctx != nil {
		if ctx.Value(field) != nil {
			v, ok := ctx.Value(field).(string)
			if ok {
				return v
			}
		}
	}
	return "NA"
}
