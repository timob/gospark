package gospark

import "github.com/timob/javabind"

type SchedulerSparkListenerExecutorAddedInterface interface {
	JavaLangObjectInterface

	// public long time()
	Time() int64

	// public java.lang.String executorId()
	ExecutorId() string

	// public org.apache.spark.scheduler.cluster.ExecutorInfo executorInfo()
	ExecutorInfo() *SchedulerClusterExecutorInfo

	// public org.apache.spark.scheduler.SparkListenerExecutorAdded copy(long, java.lang.String, org.apache.spark.scheduler.cluster.ExecutorInfo)
	Copy(a int64, b string, c SchedulerClusterExecutorInfoInterface) *SchedulerSparkListenerExecutorAdded

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type SchedulerSparkListenerExecutorAdded struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.SparkListenerExecutorAdded(long, java.lang.String, org.apache.spark.scheduler.cluster.ExecutorInfo)
func NewSchedulerSparkListenerExecutorAdded(a int64, b string, c SchedulerClusterExecutorInfoInterface) (*SchedulerSparkListenerExecutorAdded) {
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/SparkListenerExecutorAdded", a, conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("org/apache/spark/scheduler/cluster/ExecutorInfo"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &SchedulerSparkListenerExecutorAdded{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public long time()
func (jbobject *SchedulerSparkListenerExecutorAdded) Time() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "time", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public java.lang.String executorId()
func (jbobject *SchedulerSparkListenerExecutorAdded) ExecutorId() string {
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

// public org.apache.spark.scheduler.cluster.ExecutorInfo executorInfo()
func (jbobject *SchedulerSparkListenerExecutorAdded) ExecutorInfo() *SchedulerClusterExecutorInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "executorInfo", "org/apache/spark/scheduler/cluster/ExecutorInfo")
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
	unique_x := &SchedulerClusterExecutorInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.scheduler.SparkListenerExecutorAdded copy(long, java.lang.String, org.apache.spark.scheduler.cluster.ExecutorInfo)
func (jbobject *SchedulerSparkListenerExecutorAdded) Copy(a int64, b string, c SchedulerClusterExecutorInfoInterface) *SchedulerSparkListenerExecutorAdded {
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/scheduler/SparkListenerExecutorAdded", a, conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("org/apache/spark/scheduler/cluster/ExecutorInfo"))
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
	unique_x := &SchedulerSparkListenerExecutorAdded{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *SchedulerSparkListenerExecutorAdded) ProductPrefix() string {
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
func (jbobject *SchedulerSparkListenerExecutorAdded) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *SchedulerSparkListenerExecutorAdded) ProductElement(a int) *JavaLangObject {
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
func (jbobject *SchedulerSparkListenerExecutorAdded) CanEqual(a interface{}) bool {
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
func (jbobject *SchedulerSparkListenerExecutorAdded) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *SchedulerSparkListenerExecutorAdded) ToString() string {
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
func (jbobject *SchedulerSparkListenerExecutorAdded) Equals(a interface{}) bool {
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


