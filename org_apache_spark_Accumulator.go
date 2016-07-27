package gospark

import "github.com/timob/javabind"

type AccumulatorInterface interface {
	AccumulableInterface

	// public T initialValue()
	InitialValue() *JavaLangObject
}

type Accumulator struct {
	Accumulable
}

// public org.apache.spark.Accumulator(T, org.apache.spark.AccumulatorParam<T>)
func NewAccumulator(a JavaLangObjectInterface, b AccumulatorParamInterface) (*Accumulator) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/Accumulator", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/AccumulatorParam"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &Accumulator{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public T initialValue()
func (jbobject *Accumulator) InitialValue() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "initialValue", "java/lang/Object")
	if err != nil {
		panic(err)
	}
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


