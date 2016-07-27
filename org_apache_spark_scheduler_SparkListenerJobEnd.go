package gospark

import "github.com/timob/javabind"

type SchedulerSparkListenerJobEndInterface interface {
	JavaLangObjectInterface

	// public int jobId()
	JobId() int

	// public long time()
	Time() int64

	// public org.apache.spark.scheduler.JobResult jobResult()
	JobResult() *SchedulerJobResult

	// public org.apache.spark.scheduler.SparkListenerJobEnd copy(int, long, org.apache.spark.scheduler.JobResult)
	Copy(a int, b int64, c SchedulerJobResultInterface) *SchedulerSparkListenerJobEnd

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type SchedulerSparkListenerJobEnd struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.SparkListenerJobEnd(int, long, org.apache.spark.scheduler.JobResult)
func NewSchedulerSparkListenerJobEnd(a int, b int64, c SchedulerJobResultInterface) (*SchedulerSparkListenerJobEnd) {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/SparkListenerJobEnd", a, b, conv_c.Value().Cast("org/apache/spark/scheduler/JobResult"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	x := &SchedulerSparkListenerJobEnd{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int jobId()
func (jbobject *SchedulerSparkListenerJobEnd) JobId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "jobId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public long time()
func (jbobject *SchedulerSparkListenerJobEnd) Time() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "time", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.scheduler.JobResult jobResult()
func (jbobject *SchedulerSparkListenerJobEnd) JobResult() *SchedulerJobResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "jobResult", "org/apache/spark/scheduler/JobResult")
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
	unique_x := &SchedulerJobResult{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.scheduler.SparkListenerJobEnd copy(int, long, org.apache.spark.scheduler.JobResult)
func (jbobject *SchedulerSparkListenerJobEnd) Copy(a int, b int64, c SchedulerJobResultInterface) *SchedulerSparkListenerJobEnd {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/scheduler/SparkListenerJobEnd", a, b, conv_c.Value().Cast("org/apache/spark/scheduler/JobResult"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &SchedulerSparkListenerJobEnd{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *SchedulerSparkListenerJobEnd) ProductPrefix() string {
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
func (jbobject *SchedulerSparkListenerJobEnd) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *SchedulerSparkListenerJobEnd) ProductElement(a int) *JavaLangObject {
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
func (jbobject *SchedulerSparkListenerJobEnd) CanEqual(a interface{}) bool {
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
func (jbobject *SchedulerSparkListenerJobEnd) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *SchedulerSparkListenerJobEnd) ToString() string {
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
func (jbobject *SchedulerSparkListenerJobEnd) Equals(a interface{}) bool {
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


