package gospark

import "github.com/timob/javabind"

type SchedulerTaskInterface interface {
	JavaLangObjectInterface

	// public int stageId()
	StageId() int

	// public int stageAttemptId()
	StageAttemptId() int

	// public int partitionId()
	PartitionId() int

	// public void setTaskMemoryManager(org.apache.spark.memory.TaskMemoryManager)
	SetTaskMemoryManager(a MemoryTaskMemoryManagerInterface) 

	// public abstract T runTask(org.apache.spark.TaskContext)
	RunTask(a TaskContextInterface) *JavaLangObject

	// public long epoch()
	Epoch() int64

	// public org.apache.spark.TaskContextImpl context()
	Context() *TaskContextImpl

	// public long _executorDeserializeTime()
	_executorDeserializeTime() int64

	// public boolean killed()
	Killed() bool

	// public long executorDeserializeTime()
	ExecutorDeserializeTime() int64

	// public void kill(boolean)
	Kill(a bool) 
}

type SchedulerTask struct {
	JavaLangObject
}

// public int stageId()
func (jbobject *SchedulerTask) StageId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int stageAttemptId()
func (jbobject *SchedulerTask) StageAttemptId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageAttemptId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int partitionId()
func (jbobject *SchedulerTask) PartitionId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "partitionId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setTaskMemoryManager(org.apache.spark.memory.TaskMemoryManager)
func (jbobject *SchedulerTask) SetTaskMemoryManager(a MemoryTaskMemoryManagerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTaskMemoryManager", javabind.Void, conv_a.Value().Cast("org/apache/spark/memory/TaskMemoryManager"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract T runTask(org.apache.spark.TaskContext)
func (jbobject *SchedulerTask) RunTask(a TaskContextInterface) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "runTask", "java/lang/Object", conv_a.Value().Cast("org/apache/spark/TaskContext"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public long epoch()
func (jbobject *SchedulerTask) Epoch() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "epoch", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.TaskContextImpl context()
func (jbobject *SchedulerTask) Context() *TaskContextImpl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "context", "org/apache/spark/TaskContextImpl")
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
	unique_x := &TaskContextImpl{}
	unique_x.Callable = dst
	return unique_x
}

// public long _executorDeserializeTime()
func (jbobject *SchedulerTask) _executorDeserializeTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "_executorDeserializeTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public boolean killed()
func (jbobject *SchedulerTask) Killed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "killed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public long executorDeserializeTime()
func (jbobject *SchedulerTask) ExecutorDeserializeTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "executorDeserializeTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void kill(boolean)
func (jbobject *SchedulerTask) Kill(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "kill", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


