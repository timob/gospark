package gospark

import "github.com/timob/javabind"

type ApiJavaFunctionVoidFunctionInterface interface {
}

type ApiJavaFunctionVoidFunction struct {
	JavaLangObject
}

// public abstract void call(T) throws java.lang.Exception
func (jbobject *ApiJavaFunctionVoidFunction) Call(a JavaLangObjectInterface) error {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "call", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		return err
	}
	conv_a.CleanUp()
	return nil
}


