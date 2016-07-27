package gospark

import "github.com/timob/javabind"

type SchedulerSparkListenerExecutorRemovedInterface interface {
	JavaLangObjectInterface

	// public long time()
	Time() int64

	// public java.lang.String executorId()
	ExecutorId() string

	// public java.lang.String reason()
	Reason() string

	// public org.apache.spark.scheduler.SparkListenerExecutorRemoved copy(long, java.lang.String, java.lang.String)
	Copy(a int64, b string, c string) *SchedulerSparkListenerExecutorRemoved

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type SchedulerSparkListenerExecutorRemoved struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.SparkListenerExecutorRemoved(long, java.lang.String, java.lang.String)
func NewSchedulerSparkListenerExecutorRemoved(a int64, b string, c string) (*SchedulerSparkListenerExecutorRemoved) {
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/SparkListenerExecutorRemoved", a, conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &SchedulerSparkListenerExecutorRemoved{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public long time()
func (jbobject *SchedulerSparkListenerExecutorRemoved) Time() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "time", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public java.lang.String executorId()
func (jbobject *SchedulerSparkListenerExecutorRemoved) ExecutorId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "executorId", "java/lang/String")
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

// public java.lang.String reason()
func (jbobject *SchedulerSparkListenerExecutorRemoved) Reason() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "reason", "java/lang/String")
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

// public org.apache.spark.scheduler.SparkListenerExecutorRemoved copy(long, java.lang.String, java.lang.String)
func (jbobject *SchedulerSparkListenerExecutorRemoved) Copy(a int64, b string, c string) *SchedulerSparkListenerExecutorRemoved {
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/scheduler/SparkListenerExecutorRemoved", a, conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &SchedulerSparkListenerExecutorRemoved{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *SchedulerSparkListenerExecutorRemoved) ProductPrefix() string {
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
func (jbobject *SchedulerSparkListenerExecutorRemoved) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *SchedulerSparkListenerExecutorRemoved) ProductElement(a int) *JavaLangObject {
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
func (jbobject *SchedulerSparkListenerExecutorRemoved) CanEqual(a interface{}) bool {
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
func (jbobject *SchedulerSparkListenerExecutorRemoved) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *SchedulerSparkListenerExecutorRemoved) ToString() string {
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
func (jbobject *SchedulerSparkListenerExecutorRemoved) Equals(a interface{}) bool {
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


