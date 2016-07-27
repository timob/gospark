package main

import "C"

import (
	"unsafe"
	"github.com/timob/gospark"
	"github.com/timob/javabind"
	"fmt"
	"strings"
)

type funcStringBool struct {
	*gospark.ApiJavaFunctionFunctionNative
	callFunc func(string) bool	
}

func (m *funcStringBool) Call(a gospark.JavaLangObjectInterface) (*gospark.JavaLangObject, error) {
	ao := a.(*gospark.JavaLangObject)
	strObj := gospark.JavaLangString{*ao}
	str := string(strObj.GetBytes())
	v := gospark.NewJavaLangBoolean(m.callFunc(str))
	return &gospark.JavaLangObject{v.Callable}, nil
}

func newFuncStringBool(f func(string) bool) *funcStringBool {
	m := &funcStringBool{callFunc: f}
	m.ApiJavaFunctionFunctionNative = gospark.NewApiJavaFunctionFunctionNative(m)
	return m
}

//export start
func start(envPtr unsafe.Pointer) {
	javabind.SetupJVMFromEnv(envPtr)
	conf := gospark.NewSparkConf().SetAppName("Simple Application")
	inputFile := "YOUR_SPARK_HOME/README.md"
	sc := gospark.NewApiJavaJavaSparkContext2(conf)
	fileData := sc.TextFile(inputFile).Cache()
	numAs := fileData.Filter(newFuncStringBool(func(s string) bool {
		return strings.Contains(s, "a")
	})).Count()
	numBs := fileData.Filter(newFuncStringBool(func(s string) bool {
		return strings.Contains(s, "b")
	})).Count()
	fmt.Printf("Lines with a: %d, lines with b: %d\n", numAs, numBs)
} 

func init() {
}

func main() {
}
