package gospark

import "github.com/timob/javabind"

type SchedulerSparkListenerTaskStartInterface interface {
	JavaLangObjectInterface

	// public int stageId()
	StageId() int

	// public int stageAttemptId()
	StageAttemptId() int

	// public org.apache.spark.scheduler.TaskInfo taskInfo()
	TaskInfo() *SchedulerTaskInfo

	// public org.apache.spark.scheduler.SparkListenerTaskStart copy(int, int, org.apache.spark.scheduler.TaskInfo)
	Copy(a int, b int, c SchedulerTaskInfoInterface) *SchedulerSparkListenerTaskStart

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type SchedulerSparkListenerTaskStart struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.SparkListenerTaskStart(int, int, org.apache.spark.scheduler.TaskInfo)
func NewSchedulerSparkListenerTaskStart(a int, b int, c SchedulerTaskInfoInterface) (*SchedulerSparkListenerTaskStart) {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/SparkListenerTaskStart", a, b, conv_c.Value().Cast("org/apache/spark/scheduler/TaskInfo"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	x := &SchedulerSparkListenerTaskStart{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int stageId()
func (jbobject *SchedulerSparkListenerTaskStart) StageId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int stageAttemptId()
func (jbobject *SchedulerSparkListenerTaskStart) StageAttemptId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageAttemptId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.scheduler.TaskInfo taskInfo()
func (jbobject *SchedulerSparkListenerTaskStart) TaskInfo() *SchedulerTaskInfo {
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

// public org.apache.spark.scheduler.SparkListenerTaskStart copy(int, int, org.apache.spark.scheduler.TaskInfo)
func (jbobject *SchedulerSparkListenerTaskStart) Copy(a int, b int, c SchedulerTaskInfoInterface) *SchedulerSparkListenerTaskStart {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/scheduler/SparkListenerTaskStart", a, b, conv_c.Value().Cast("org/apache/spark/scheduler/TaskInfo"))
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
	unique_x := &SchedulerSparkListenerTaskStart{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *SchedulerSparkListenerTaskStart) ProductPrefix() string {
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
func (jbobject *SchedulerSparkListenerTaskStart) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *SchedulerSparkListenerTaskStart) ProductElement(a int) *JavaLangObject {
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
func (jbobject *SchedulerSparkListenerTaskStart) CanEqual(a interface{}) bool {
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
func (jbobject *SchedulerSparkListenerTaskStart) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *SchedulerSparkListenerTaskStart) ToString() string {
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
func (jbobject *SchedulerSparkListenerTaskStart) Equals(a interface{}) bool {
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


