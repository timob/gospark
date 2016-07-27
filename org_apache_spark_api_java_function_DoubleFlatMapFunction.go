package gospark

import "github.com/timob/javabind"

type ApiJavaFunctionDoubleFlatMapFunctionInterface interface {
}

type ApiJavaFunctionDoubleFlatMapFunction struct {
	JavaLangObject
}

// public abstract java.lang.Iterable<java.lang.Double> call(T) throws java.lang.Exception
func (jbobject *ApiJavaFunctionDoubleFlatMapFunction) Call(a JavaLangObjectInterface) (*JavaLangIterable, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "call", "java/lang/Iterable", conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		var zero *JavaLangIterable
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
	unique_x := &JavaLangIterable{}
	unique_x.Callable = dst
	return unique_x, nil
}


