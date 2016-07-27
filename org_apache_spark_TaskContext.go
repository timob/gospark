package gospark

import "github.com/timob/javabind"

type TaskContextInterface interface {
	JavaLangObjectInterface

	// public static int getPartitionId()
	GetPartitionId() int

	// public static org.apache.spark.TaskContext get()
	Get() *TaskContext

	// public abstract boolean isCompleted()
	IsCompleted() bool

	// public abstract boolean isInterrupted()
	IsInterrupted() bool

	// public abstract boolean runningLocally()
	RunningLocally() bool

	// public abstract boolean isRunningLocally()
	IsRunningLocally() bool

	// public abstract org.apache.spark.TaskContext addTaskCompletionListener(org.apache.spark.util.TaskCompletionListener)
	AddTaskCompletionListener(a UtilTaskCompletionListenerInterface) *TaskContext

	// public abstract org.apache.spark.TaskContext addTaskFailureListener(org.apache.spark.util.TaskFailureListener)
	AddTaskFailureListener(a UtilTaskFailureListenerInterface) *TaskContext

	// public abstract int stageId()
	StageId() int

	// public abstract int partitionId()
	PartitionId() int

	// public abstract int attemptNumber()
	AttemptNumber() int

	// public abstract long attemptId()
	AttemptId() int64

	// public abstract long taskAttemptId()
	TaskAttemptId() int64

	// public abstract org.apache.spark.executor.TaskMetrics taskMetrics()
	TaskMetrics() *ExecutorTaskMetrics

	// public abstract org.apache.spark.memory.TaskMemoryManager taskMemoryManager()
	TaskMemoryManager() *MemoryTaskMemoryManager

	// public abstract void registerAccumulator(org.apache.spark.Accumulable<?, ?>)
	RegisterAccumulator(a AccumulableInterface) 
}

type TaskContext struct {
	JavaLangObject
}

// public org.apache.spark.TaskContext()
func NewTaskContext() (*TaskContext) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/TaskContext")
	if err != nil {
		panic(err)
	}
	x := &TaskContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static int getPartitionId()
func (jbobject *TaskContext) GetPartitionId() int {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/TaskContext", "getPartitionId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public static org.apache.spark.TaskContext get()
func (jbobject *TaskContext) Get() *TaskContext {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/TaskContext", "get", "org/apache/spark/TaskContext")
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
	unique_x := &TaskContext{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract boolean isCompleted()
func (jbobject *TaskContext) IsCompleted() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isCompleted", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract boolean isInterrupted()
func (jbobject *TaskContext) IsInterrupted() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isInterrupted", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract boolean runningLocally()
func (jbobject *TaskContext) RunningLocally() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "runningLocally", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract boolean isRunningLocally()
func (jbobject *TaskContext) IsRunningLocally() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isRunningLocally", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract org.apache.spark.TaskContext addTaskCompletionListener(org.apache.spark.util.TaskCompletionListener)
func (jbobject *TaskContext) AddTaskCompletionListener(a UtilTaskCompletionListenerInterface) *TaskContext {
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

// public abstract org.apache.spark.TaskContext addTaskFailureListener(org.apache.spark.util.TaskFailureListener)
func (jbobject *TaskContext) AddTaskFailureListener(a UtilTaskFailureListenerInterface) *TaskContext {
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

// public abstract int stageId()
func (jbobject *TaskContext) StageId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int partitionId()
func (jbobject *TaskContext) PartitionId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "partitionId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int attemptNumber()
func (jbobject *TaskContext) AttemptNumber() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "attemptNumber", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract long attemptId()
func (jbobject *TaskContext) AttemptId() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "attemptId", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public abstract long taskAttemptId()
func (jbobject *TaskContext) TaskAttemptId() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "taskAttemptId", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public abstract org.apache.spark.executor.TaskMetrics taskMetrics()
func (jbobject *TaskContext) TaskMetrics() *ExecutorTaskMetrics {
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

// public abstract org.apache.spark.memory.TaskMemoryManager taskMemoryManager()
func (jbobject *TaskContext) TaskMemoryManager() *MemoryTaskMemoryManager {
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

// public abstract void registerAccumulator(org.apache.spark.Accumulable<?, ?>)
func (jbobject *TaskContext) RegisterAccumulator(a AccumulableInterface)  {
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


