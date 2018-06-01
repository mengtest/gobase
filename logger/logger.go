/**************************************************************************************
Code Description    : 日志
Code Vesion         :
					|------------------------------------------------------------|
						  Version    					Editor            Time
							1.0        					yuansudong        2016.4.12
					|------------------------------------------------------------|
Version Description	:
                    |------------------------------------------------------------|
						  Version
							1.0
								 ....
					|------------------------------------------------------------|
***************************************************************************************/

package logger

import (
	"fmt"
	"runtime"
	"time"
)

var (
	logHandle *log
	dbg       string
	inf       string
	err       string
)

const (
	// LogDebugLevel Debug级别
	LogDebugLevel = 0
	// LogInfoLevel Info 级别
	LogInfoLevel = 1
	// LogErrLevel Error 级别
	LogErrLevel = 2
)

func init() {
	logHandle = &log{
		level: LogDebugLevel,
	}
	inf = "[INF]"
	dbg = "[DBG]"
	err = "[ERR]"
}

// log
type log struct {
	level int
}

//  SetLogLevel 设置日志的等级
func (l *log) SetLogLevel(level int) {
	l.level = level
}

// Debug 用于打印调试信息
func (l *log) debug(stack, format string, args ...interface{}) {
	if LogDebugLevel >= l.level {
		args1 := []interface{}{
			getDate(),
			"   ",
			dbg,
			"   ",
			stack,
			"   ",
			fmt.Sprintf(format, args...),
		}
		fmt.Println(
			args1...,
		)
	}
}

// Info 用于打印调试信息
func (l *log) info(stack, format string, args ...interface{}) {
	if LogInfoLevel >= l.level {
		args1 := []interface{}{
			getDate(),
			"   ",
			inf,
			"   ",
			stack,
			"   ",
			fmt.Sprintf(format, args...),
		}
		fmt.Println(
			args1...,
		)
	}
}

// Info 用于打印调试信息
func (l *log) error(stack, format string, args ...interface{}) {
	if LogErrLevel >= l.level {
		args1 := []interface{}{
			getDate(),
			"   ",
			err,
			"   ",
			stack,
			"   ",
			fmt.Sprintf(format, args...),
		}
		fmt.Println(
			args1...,
		)
	}
}

// Debug 用于打印Debug类型的消息
func Debug(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	logHandle.debug(fmt.Sprintf("%s:%d", file, line), format, args...)
}

// Info 用于打印Info级别的消息
func Info(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	logHandle.info(fmt.Sprintf("%s:%d", file, line), format, args...)
}

// Error 用于打印错误级别的消息
func Error(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	logHandle.error(fmt.Sprintf("%s:%d", file, line), format, args...)
}

func getDate() string {
	t := time.Now()
	return fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
}
