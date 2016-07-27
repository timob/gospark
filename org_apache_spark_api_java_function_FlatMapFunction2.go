package gospark

import "github.com/timob/javabind"

type ApiJavaFunctionFlatMapFunction2Interface interface {
}

type ApiJavaFunctionFlatMapFunction2 struct {
	JavaLangObject
}

// public abstract java.lang.Iterable<R> call(T1, T2) throws java.lang.Exception
func (jbobject *ApiJavaFunctionFlatMapFunction2) Call(a JavaLangObjectInterface, b JavaLangObjectInterface) (*JavaLangIterable, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "call", "java/lang/Iterable", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		var zero *JavaLangIterable
		return zero, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
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


