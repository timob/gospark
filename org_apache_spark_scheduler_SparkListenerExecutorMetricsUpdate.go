package gospark

import "github.com/timob/javabind"

type SchedulerSparkListenerExecutorMetricsUpdateInterface interface {
	JavaLangObjectInterface

	// public java.lang.String execId()
	ExecId() string

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type SchedulerSparkListenerExecutorMetricsUpdate struct {
	JavaLangObject
}

// public java.lang.String execId()
func (jbobject *SchedulerSparkListenerExecutorMetricsUpdate) ExecId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "execId", "java/lang/String")
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

// public java.lang.String productPrefix()
func (jbobject *SchedulerSparkListenerExecutorMetricsUpdate) ProductPrefix() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productPrefix", "java/lang/String")
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

// public int productArity()
func (jbobject *SchedulerSparkListenerExecutorMetricsUpdate) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *SchedulerSparkListenerExecutorMetricsUpdate) ProductElement(a int) *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productElement", "java/lang/Object", a)
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

// public boolean canEqual(java.lang.Object)
func (jbobject *SchedulerSparkListenerExecutorMetricsUpdate) CanEqual(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "canEqual", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *SchedulerSparkListenerExecutorMetricsUpdate) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *SchedulerSparkListenerExecutorMetricsUpdate) ToString() string {
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

// public boolean equals(java.lang.Object)
func (jbobject *SchedulerSparkListenerExecutorMetricsUpdate) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}


