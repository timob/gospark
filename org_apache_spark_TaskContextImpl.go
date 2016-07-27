package gospark

import "github.com/timob/javabind"

type TaskContextImplInterface interface {
	TaskContextInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public org.apache.spark.TaskContextImpl addTaskCompletionListener(org.apache.spark.util.TaskCompletionListener)
	AddTaskCompletionListener2(a UtilTaskCompletionListenerInterface) *TaskContextImpl

	// public org.apache.spark.TaskContextImpl addTaskFailureListener(org.apache.spark.util.TaskFailureListener)
	AddTaskFailureListener2(a UtilTaskFailureListenerInterface) *TaskContextImpl

	// public void markTaskCompleted()
	MarkTaskCompleted() 

	// public void markInterrupted()
	MarkInterrupted() 
}

type TaskContextImpl struct {
	TaskContext
}

// public java.lang.String logName()
func (jbobject *TaskContextImpl) LogName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "logName", "java/lang/String")
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

// public boolean isTraceEnabled()
func (jbobject *TaskContextImpl) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int stageId()
func (jbobject *TaskContextImpl) StageId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int partitionId()
func (jbobject *TaskContextImpl) PartitionId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "partitionId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public long taskAttemptId()
func (jbobject *TaskContextImpl) TaskAttemptId() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "taskAttemptId", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public int attemptNumber()
func (jbobject *TaskContextImpl) AttemptNumber() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "attemptNumber", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.memory.TaskMemoryManager taskMemoryManager()
func (jbobject *TaskContextImpl) TaskMemoryManager() *MemoryTaskMemoryManager {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "taskMemoryManager", "org/apache/spark/memory/TaskMemoryManager")
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
	unique_x := &MemoryTaskMemoryManager{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean runningLocally()
func (jbobject *TaskContextImpl) RunningLocally() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "runningLocally", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.executor.TaskMetrics taskMetrics()
func (jbobject *TaskContextImpl) TaskMetrics() *ExecutorTaskMetrics {
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

// public long attemptId()
func (jbobject *TaskContextImpl) AttemptId() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "attemptId", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.TaskContextImpl addTaskCompletionListener(org.apache.spark.util.TaskCompletionListener)
func (jbobject *TaskContextImpl) AddTaskCompletionListener2(a UtilTaskCompletionListenerInterface) *TaskContextImpl {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "addTaskCompletionListener", "org/apache/spark/TaskContextImpl", conv_a.Value().Cast("org/apache/spark/util/TaskCompletionListener"))
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
	unique_x := &TaskContextImpl{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.TaskContextImpl addTaskFailureListener(org.apache.spark.util.TaskFailureListener)
func (jbobject *TaskContextImpl) AddTaskFailureListener2(a UtilTaskFailureListenerInterface) *TaskContextImpl {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "addTaskFailureListener", "org/apache/spark/TaskContextImpl", conv_a.Value().Cast("org/apache/spark/util/TaskFailureListener"))
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
	unique_x := &TaskContextImpl{}
	unique_x.Callable = dst
	return unique_x
}

// public void markTaskCompleted()
func (jbobject *TaskContextImpl) MarkTaskCompleted()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "markTaskCompleted", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void markInterrupted()
func (jbobject *TaskContextImpl) MarkInterrupted()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "markInterrupted", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public boolean isCompleted()
func (jbobject *TaskContextImpl) IsCompleted() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isCompleted", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isRunningLocally()
func (jbobject *TaskContextImpl) IsRunningLocally() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isRunningLocally", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isInterrupted()
func (jbobject *TaskContextImpl) IsInterrupted() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isInterrupted", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public synchronized void registerAccumulator(org.apache.spark.Accumulable<?, ?>)
func (jbobject *TaskContextImpl) RegisterAccumulator(a AccumulableInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "registerAccumulator", javabind.Void, conv_a.Value().Cast("org/apache/spark/Accumulable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.apache.spark.TaskContext addTaskFailureListener(org.apache.spark.util.TaskFailureListener)
func (jbobject *TaskContextImpl) AddTaskFailureListener(a UtilTaskFailureListenerInterface) *TaskContext {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "addTaskFailureListener", "org/apache/spark/TaskContext", conv_a.Value().Cast("org/apache/spark/util/TaskFailureListener"))
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
	unique_x := &TaskContext{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.TaskContext addTaskCompletionListener(org.apache.spark.util.TaskCompletionListener)
func (jbobject *TaskContextImpl) AddTaskCompletionListener(a UtilTaskCompletionListenerInterface) *TaskContext {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "addTaskCompletionListener", "org/apache/spark/TaskContext", conv_a.Value().Cast("org/apache/spark/util/TaskCompletionListener"))
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
	unique_x := &TaskContext{}
	unique_x.Callable = dst
	return unique_x
}


