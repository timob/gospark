package gospark

import "github.com/timob/javabind"

type ApiJavaFunctionDoubleFunctionInterface interface {
}

type ApiJavaFunctionDoubleFunction struct {
	JavaLangObject
}

// public abstract double call(T) throws java.lang.Exception
func (jbobject *ApiJavaFunctionDoubleFunction) Call(a JavaLangObjectInterface) (float64, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "call", javabind.Double, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		var zero float64
		return zero, err
	}
	conv_a.CleanUp()
	return jret.(float64), nil
}


