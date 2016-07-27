package gospark

import "github.com/timob/javabind"

type PartialPartialResultInterface interface {
	JavaLangObjectInterface

	// public R initialValue()
	InitialValue() *JavaLangObject

	// public boolean isInitialValueFinal()
	IsInitialValueFinal() bool

	// public synchronized R getFinalValue()
	GetFinalValue() *JavaLangObject

	// public synchronized void setFinalValue(R)
	SetFinalValue(a JavaLangObjectInterface) 
}

type PartialPartialResult struct {
	JavaLangObject
}

// public org.apache.spark.partial.PartialResult(R, boolean)
func NewPartialPartialResult(a JavaLangObjectInterface, b bool) (*PartialPartialResult) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/partial/PartialResult", conv_a.Value().Cast("java/lang/Object"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &PartialPartialResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public R initialValue()
func (jbobject *PartialPartialResult) InitialValue() *JavaLangObject {
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

// public boolean isInitialValueFinal()
func (jbobject *PartialPartialResult) IsInitialValueFinal() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isInitialValueFinal", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public synchronized R getFinalValue()
func (jbobject *PartialPartialResult) GetFinalValue() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFinalValue", "java/lang/Object")
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

// public synchronized void setFinalValue(R)
func (jbobject *PartialPartialResult) SetFinalValue(a JavaLangObjectInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFinalValue", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized java.lang.String toString()
func (jbobject *PartialPartialResult) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}


