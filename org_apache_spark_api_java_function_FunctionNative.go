package gospark

import "github.com/timob/javabind"
import "unsafe"
import "sync"


/*
#include<stdint.h>

extern uintptr_t go_callback_ApiJavaFunctionFunctionNative_Call(void *env, uintptr_t obj, uintptr_t arg_0 );

*/
import "C"
//export go_callback_ApiJavaFunctionFunctionNative_Call
func go_callback_ApiJavaFunctionFunctionNative_Call(env unsafe.Pointer, obj uintptr, arg_0 uintptr) uintptr {
	javabind.AddEnv(env)
    rObj := &javabind.Callable{javabind.WrapJObject(obj, "org/apache/spark/api/java/function/FunctionNative", false)}
    id, err := rObj.GetField(javabind.GetEnv(), "id", javabind.Int)
    if err != nil {
        panic(err)
    }

    i := ApiJavaFunctionFunctionNativeMap[id.(int)]
    	retcon_a := javabind.NewJavaToGoCallable()
	dst_a := &javabind.Callable{}
	retcon_a.Dest(dst_a)
	if err := retcon_a.Convert(javabind.WrapJObject(arg_0, "java/lang/Object", false)); err != nil {
		panic(err)
	}
	arg_a := &JavaLangObject{}
	arg_a.Callable = dst_a
ret, _ := i.(ApiJavaFunctionFunctionExtraInterface).Call(arg_a)
	retconv := javabind.NewGoToJavaCallable()
	if err := retconv.Convert(ret); err != nil {
	panic(err)
	}
	return uintptr(retconv.Value().JObject())
}

var ApiJavaFunctionFunctionNativeMap = make(map[int]ApiJavaFunctionFunctionExtraInterface)
var ApiJavaFunctionFunctionNativeCount int = 0
var ApiJavaFunctionFunctionNativeMutex sync.Mutex

type ApiJavaFunctionFunctionNative struct {
	*javabind.Callable
	ApiJavaFunctionFunctionInterface
}

func NewApiJavaFunctionFunctionNative(implementation ApiJavaFunctionFunctionExtraInterface) *ApiJavaFunctionFunctionNative {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/function/FunctionNative")
	if err != nil {
		panic(err)
	}
	
	x := &ApiJavaFunctionFunctionNative{}
	x.Callable = &javabind.Callable{obj}
	ApiJavaFunctionFunctionNativeMutex.Lock()
	defer ApiJavaFunctionFunctionNativeMutex.Unlock()
	ApiJavaFunctionFunctionNativeCount++

    err = x.Callable.SetField(javabind.GetEnv(), "id", ApiJavaFunctionFunctionNativeCount)
    if err != nil {
        panic(err)
    }
    ApiJavaFunctionFunctionNativeMap[ApiJavaFunctionFunctionNativeCount] = implementation
	
	return x
}


    func init() {
        javabind.OnJVMStart(func() {
        javabind.GetEnv().RegisterNative("org/apache/spark/api/java/function/FunctionNative", "call", "java/lang/Object", []interface{}{"java/lang/Object"}, C.go_callback_ApiJavaFunctionFunctionNative_Call)

        })
    }
    type ApiJavaFunctionFunctionExtraInterface interface {

	// public abstract R call(T1) throws java.lang.Exception
	Call(a JavaLangObjectInterface) (*JavaLangObject, error)
}

