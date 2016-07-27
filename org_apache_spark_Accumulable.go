package gospark

import "github.com/timob/javabind"

type AccumulableInterface interface {
	JavaLangObjectInterface

	// public long id()
	Id() int64

	// public R zero()
	Zero() *JavaLangObject

	// public boolean isInternal()
	IsInternal() bool

	// public void add(T)
	Add(a JavaLangObjectInterface) 

	// public void merge(R)
	Merge(a JavaLangObjectInterface) 

	// public R value()
	Value() *JavaLangObject

	// public R localValue()
	LocalValue() *JavaLangObject

	// public void setValue(R)
	SetValue(a JavaLangObjectInterface) 
}

type Accumulable struct {
	JavaLangObject
}

// public org.apache.spark.Accumulable(R, org.apache.spark.AccumulableParam<R, T>, boolean)
func NewAccumulable2(a JavaLangObjectInterface, b AccumulableParamInterface, c bool) (*Accumulable) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/Accumulable", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/AccumulableParam"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &Accumulable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.Accumulable(R, org.apache.spark.AccumulableParam<R, T>)
func NewAccumulable(a JavaLangObjectInterface, b AccumulableParamInterface) (*Accumulable) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/Accumulable", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/AccumulableParam"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &Accumulable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public long id()
func (jbobject *Accumulable) Id() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "id", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public R zero()
func (jbobject *Accumulable) Zero() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "zero", "java/lang/Object")
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

// public boolean isInternal()
func (jbobject *Accumulable) IsInternal() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isInternal", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void add(T)
func (jbobject *Accumulable) Add(a JavaLangObjectInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "add", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void merge(R)
func (jbobject *Accumulable) Merge(a JavaLangObjectInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "merge", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public R value()
func (jbobject *Accumulable) Value() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "value", "java/lang/Object")
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

// public R localValue()
func (jbobject *Accumulable) LocalValue() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "localValue", "java/lang/Object")
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

// public void setValue(R)
func (jbobject *Accumulable) SetValue(a JavaLangObjectInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValue", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String toString()
func (jbobject *Accumulable) ToString() string {
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


