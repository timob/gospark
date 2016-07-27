package gospark

import "github.com/timob/javabind"

type ApiJavaFunctionFunctionInterface interface {
}

type ApiJavaFunctionFunction struct {
	JavaLangObject
}

// public abstract R call(T1) throws java.lang.Exception
func (jbobject *ApiJavaFunctionFunction) Call(a JavaLangObjectInterface) (*JavaLangObject, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "call", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		var zero *JavaLangObject
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x, nil
}


