package utils

import (
	"log"
	"runtime"
	"strings"
)

func CallerFunctionName(callback int) string {
	pc, _, _, ok := runtime.Caller(0 + callback)
	if !ok {
		return "unknown"
	}
	callerFunction := runtime.FuncForPC(pc)
	if callerFunction == nil {
		return "unknown"
	}
	functionName := callerFunction.Name()
	functionName = functionName[strings.LastIndex(functionName, ".")+1:]
	return functionName
}

func Log(message string) {
	log.Printf("%s: %s", CallerFunctionName(2), message)
}
