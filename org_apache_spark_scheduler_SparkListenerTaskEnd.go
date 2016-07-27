package gospark

import "github.com/timob/javabind"

type SchedulerSparkListenerTaskEndInterface interface {
	JavaLangObjectInterface

	// public int stageId()
	StageId() int

	// public int stageAttemptId()
	StageAttemptId() int

	// public java.lang.String taskType()
	TaskType() string

	// public org.apache.spark.TaskEndReason reason()
	Reason() *TaskEndReason

	// public org.apache.spark.scheduler.TaskInfo taskInfo()
	TaskInfo() *SchedulerTaskInfo

	// public org.apache.spark.executor.TaskMetrics taskMetrics()
	TaskMetrics() *ExecutorTaskMetrics

	// public org.apache.spark.scheduler.SparkListenerTaskEnd copy(int, int, java.lang.String, org.apache.spark.TaskEndReason, org.apache.spark.scheduler.TaskInfo, org.apache.spark.executor.TaskMetrics)
	Copy(a int, b int, c string, d TaskEndReasonInterface, e SchedulerTaskInfoInterface, f ExecutorTaskMetricsInterface) *SchedulerSparkListenerTaskEnd

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type SchedulerSparkListenerTaskEnd struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.SparkListenerTaskEnd(int, int, java.lang.String, org.apache.spark.TaskEndReason, org.apache.spark.scheduler.TaskInfo, org.apache.spark.executor.TaskMetrics)
func NewSchedulerSparkListenerTaskEnd(a int, b int, c string, d TaskEndReasonInterface, e SchedulerTaskInfoInterface, f ExecutorTaskMetricsInterface) (*SchedulerSparkListenerTaskEnd) {
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaCallable()
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/SparkListenerTaskEnd", a, b, conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("org/apache/spark/TaskEndReason"), conv_e.Value().Cast("org/apache/spark/scheduler/TaskInfo"), conv_f.Value().Cast("org/apache/spark/executor/TaskMetrics"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &SchedulerSparkListenerTaskEnd{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int stageId()
func (jbobject *SchedulerSparkListenerTaskEnd) StageId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int stageAttemptId()
func (jbobject *SchedulerSparkListenerTaskEnd) StageAttemptId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageAttemptId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String taskType()
func (jbobject *SchedulerSparkListenerTaskEnd) TaskType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "taskType", "java/lang/String")
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

// public org.apache.spark.TaskEndReason reason()
func (jbobject *SchedulerSparkListenerTaskEnd) Reason() *TaskEndReason {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "reason", "org/apache/spark/TaskEndReason")
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
	unique_x := &TaskEndReason{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.scheduler.TaskInfo taskInfo()
func (jbobject *SchedulerSparkListenerTaskEnd) TaskInfo() *SchedulerTaskInfo {
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

// public org.apache.spark.executor.TaskMetrics taskMetrics()
func (jbobject *SchedulerSparkListenerTaskEnd) TaskMetrics() *ExecutorTaskMetrics {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "taskMetrics", "org/apache/spark/executor/TaskMetrics")
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
	unique_x := &ExecutorTaskMetrics{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.scheduler.SparkListenerTaskEnd copy(int, int, java.lang.String, org.apache.spark.TaskEndReason, org.apache.spark.scheduler.TaskInfo, org.apache.spark.executor.TaskMetrics)
func (jbobject *SchedulerSparkListenerTaskEnd) Copy(a int, b int, c string, d TaskEndReasonInterface, e SchedulerTaskInfoInterface, f ExecutorTaskMetricsInterface) *SchedulerSparkListenerTaskEnd {
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaCallable()
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/scheduler/SparkListenerTaskEnd", a, b, conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("org/apache/spark/TaskEndReason"), conv_e.Value().Cast("org/apache/spark/scheduler/TaskInfo"), conv_f.Value().Cast("org/apache/spark/executor/TaskMetrics"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	conv_f.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &SchedulerSparkListenerTaskEnd{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *SchedulerSparkListenerTaskEnd) ProductPrefix() string {
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
func (jbobject *SchedulerSparkListenerTaskEnd) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *SchedulerSparkListenerTaskEnd) ProductElement(a int) *JavaLangObject {
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
func (jbobject *SchedulerSparkListenerTaskEnd) CanEqual(a interface{}) bool {
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
func (jbobject *SchedulerSparkListenerTaskEnd) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *SchedulerSparkListenerTaskEnd) ToString() string {
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
func (jbobject *SchedulerSparkListenerTaskEnd) Equals(a interface{}) bool {
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


