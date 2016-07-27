package gospark

import "github.com/timob/javabind"

type AccumulatorParamInterface interface {

	// public abstract T addAccumulator(T, T)
	AddAccumulator(a JavaLangObjectInterface, b JavaLangObjectInterface) *JavaLangObject
}

type AccumulatorParam struct {
	JavaLangObject
}

// public abstract T addAccumulator(T, T)
func (jbobject *AccumulatorParam) AddAccumulator(a JavaLangObjectInterface, b JavaLangObjectInterface) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "addAccumulator", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}


