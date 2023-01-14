package glog

import (
	"fmt"
	"time"
)

var (
	// black   = color("\033[1;30m%s\033[0m")
	red    = color("\033[1;31m%s\033[0m")
	green  = color("\033[1;32m%s\033[0m")
	yellow = color("\033[1;33m%s\033[0m")
	// purple  = color("\033[1;34m%s\033[0m")
	// magenta = color("\033[1;35m%s\033[0m")
	teal = color("\033[1;36m%s\033[0m")
	// white   = color("\033[1;37m%s\033[0m")
)

// Изменение цвета в консоли
func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args))
	}
	return sprint
}

// логирование уровня Success
func (gl *GLog) Success(content interface{}) {
	infoContent := fmt.Sprintf("%+v %v", content, time.Now().Format(time.RFC3339))
	fmt.Print(green("[success] "))
	fmt.Println(infoContent)
	gl.loggerWriter(infoContent, "success")
}

// логирование уровня Log
func (gl *GLog) Log(content interface{}) {
	infoContent := fmt.Sprintf("%+v %v", content, time.Now().Format(time.RFC3339))
	fmt.Print(teal("[info] "))
	fmt.Println(infoContent)
	gl.loggerWriter(infoContent, "info")
}

// логирование уровня Warn
func (gl *GLog) Warn(content interface{}) {
	infoContent := fmt.Sprintf("%+v %v", content, time.Now().Format(time.RFC3339))
	fmt.Print(yellow("[warn] "))
	fmt.Println(infoContent)
	gl.loggerWriter(infoContent, "warn")
}

// логирование уровня Error
func (gl *GLog) Error(content interface{}, err error) {
	infoContent := fmt.Sprintf("%+v %v", content, time.Now().Format(time.RFC3339))
	fmt.Print(red("[error] "))

	trace := getTrace(infoContent, err)
	fmt.Println(infoContent)
	fmt.Print("trace: ")
	printTrace := trace
	if len(trace) > 4 {
		printTrace = trace[2:4]
	}
	for _, v := range printTrace {
		fmt.Println(v)
	}

	gl.loggerWriter(infoContent, "error")
}
