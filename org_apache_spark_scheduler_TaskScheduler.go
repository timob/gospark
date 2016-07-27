package gospark

import "github.com/timob/javabind"

type SchedulerTaskSchedulerInterface interface {

	// public abstract org.apache.spark.scheduler.Pool rootPool()
	RootPool() *SchedulerPool

	// public abstract void start()
	Start() 

	// public abstract void postStartHook()
	PostStartHook() 

	// public abstract void stop()
	Stop() 

	// public abstract void submitTasks(org.apache.spark.scheduler.TaskSet)
	SubmitTasks(a SchedulerTaskSetInterface) 

	// public abstract void cancelTasks(int, boolean)
	CancelTasks(a int, b bool) 

	// public abstract void setDAGScheduler(org.apache.spark.scheduler.DAGScheduler)
	SetDAGScheduler(a SchedulerDAGSchedulerInterface) 

	// public abstract int defaultParallelism()
	DefaultParallelism() int

	// public abstract java.lang.String applicationId()
	ApplicationId() string

	// public abstract void executorLost(java.lang.String, org.apache.spark.scheduler.ExecutorLossReason)
	ExecutorLost(a string, b SchedulerExecutorLossReasonInterface) 
}

type SchedulerTaskScheduler struct {
	JavaLangObject
}

// public abstract org.apache.spark.scheduler.Pool rootPool()
func (jbobject *SchedulerTaskScheduler) RootPool() *SchedulerPool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "rootPool", "org/apache/spark/scheduler/Pool")
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
	unique_x := &SchedulerPool{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract void start()
func (jbobject *SchedulerTaskScheduler) Start()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "start", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract void postStartHook()
func (jbobject *SchedulerTaskScheduler) PostStartHook()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "postStartHook", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract void stop()
func (jbobject *SchedulerTaskScheduler) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract void submitTasks(org.apache.spark.scheduler.TaskSet)
func (jbobject *SchedulerTaskScheduler) SubmitTasks(a SchedulerTaskSetInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "submitTasks", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/TaskSet"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void cancelTasks(int, boolean)
func (jbobject *SchedulerTaskScheduler) CancelTasks(a int, b bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelTasks", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public abstract void setDAGScheduler(org.apache.spark.scheduler.DAGScheduler)
func (jbobject *SchedulerTaskScheduler) SetDAGScheduler(a SchedulerDAGSchedulerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDAGScheduler", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/DAGScheduler"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract int defaultParallelism()
func (jbobject *SchedulerTaskScheduler) DefaultParallelism() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "defaultParallelism", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract java.lang.String applicationId()
func (jbobject *SchedulerTaskScheduler) ApplicationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "applicationId", "java/lang/String")
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

// public abstract void executorLost(java.lang.String, org.apache.spark.scheduler.ExecutorLossReason)
func (jbobject *SchedulerTaskScheduler) ExecutorLost(a string, b SchedulerExecutorLossReasonInterface)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "executorLost", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("org/apache/spark/scheduler/ExecutorLossReason"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}


