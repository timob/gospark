package gospark

import "github.com/timob/javabind"

type SchedulerSparkListenerTaskGettingResultInterface interface {
	JavaLangObjectInterface

	// public org.apache.spark.scheduler.TaskInfo taskInfo()
	TaskInfo() *SchedulerTaskInfo

	// public org.apache.spark.scheduler.SparkListenerTaskGettingResult copy(org.apache.spark.scheduler.TaskInfo)
	Copy(a SchedulerTaskInfoInterface) *SchedulerSparkListenerTaskGettingResult

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type SchedulerSparkListenerTaskGettingResult struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.SparkListenerTaskGettingResult(org.apache.spark.scheduler.TaskInfo)
func NewSchedulerSparkListenerTaskGettingResult(a SchedulerTaskInfoInterface) (*SchedulerSparkListenerTaskGettingResult) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/SparkListenerTaskGettingResult", conv_a.Value().Cast("org/apache/spark/scheduler/TaskInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &SchedulerSparkListenerTaskGettingResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.scheduler.TaskInfo taskInfo()
func (jbobject *SchedulerSparkListenerTaskGettingResult) TaskInfo() *SchedulerTaskInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "taskInfo", "org/apache/spark/scheduler/TaskInfo")
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
	unique_x := &SchedulerTaskInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.scheduler.SparkListenerTaskGettingResult copy(org.apache.spark.scheduler.TaskInfo)
func (jbobject *SchedulerSparkListenerTaskGettingResult) Copy(a SchedulerTaskInfoInterface) *SchedulerSparkListenerTaskGettingResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/scheduler/SparkListenerTaskGettingResult", conv_a.Value().Cast("org/apache/spark/scheduler/TaskInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &SchedulerSparkListenerTaskGettingResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *SchedulerSparkListenerTaskGettingResult) ProductPrefix() string {
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
func (jbobject *SchedulerSparkListenerTaskGettingResult) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *SchedulerSparkListenerTaskGettingResult) ProductElement(a int) *JavaLangObject {
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
func (jbobject *SchedulerSparkListenerTaskGettingResult) CanEqual(a interface{}) bool {
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
func (jbobject *SchedulerSparkListenerTaskGettingResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *SchedulerSparkListenerTaskGettingResult) ToString() string {
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
func (jbobject *SchedulerSparkListenerTaskGettingResult) Equals(a interface{}) bool {
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


