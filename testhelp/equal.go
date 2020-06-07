package testhelp

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func Equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal(strings.Join(StackTrace(), "\n\t\t\t"), "Expected:\n", expected, "\nGot:\n", actual)
	}
}

func StackTrace() []string {
	callStack := []string{}
	for i := 0; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			return callStack[2:]
		}

		funcPackageName := runtime.FuncForPC(pc).Name()
		if strings.HasPrefix(funcPackageName, "testing.") {
			return callStack[2:]
		}

		funcParts := strings.Split(funcPackageName, "/")
		funcName := funcParts[len(funcParts)-1]

		fileParts := strings.Split(file, "/")
		fileName := fileParts[len(fileParts)-1]

		callStack = append(callStack, fmt.Sprintf("%s(%s:%d)", fileName, funcName, line))
	}
}
