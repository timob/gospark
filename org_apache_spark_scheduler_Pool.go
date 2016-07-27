package gospark

import "github.com/timob/javabind"

type SchedulerPoolInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public java.lang.String poolName()
	PoolName() string

	// public int weight()
	Weight() int

	// public int minShare()
	MinShare() int

	// public int runningTasks()
	RunningTasks() int

	// public int priority()
	Priority() int

	// public int stageId()
	StageId() int

	// public java.lang.String name()
	Name() string

	// public org.apache.spark.scheduler.Pool parent()
	Parent() *SchedulerPool

	// public org.apache.spark.scheduler.SchedulingAlgorithm taskSetSchedulingAlgorithm()
	TaskSetSchedulingAlgorithm() *SchedulerSchedulingAlgorithm

	// public void addSchedulable(org.apache.spark.scheduler.Schedulable)
	AddSchedulable(a SchedulerSchedulableInterface) 

	// public void removeSchedulable(org.apache.spark.scheduler.Schedulable)
	RemoveSchedulable(a SchedulerSchedulableInterface) 

	// public org.apache.spark.scheduler.Schedulable getSchedulableByName(java.lang.String)
	GetSchedulableByName(a string) *SchedulerSchedulable

	// public void executorLost(java.lang.String, java.lang.String, org.apache.spark.scheduler.ExecutorLossReason)
	ExecutorLost(a string, b string, c SchedulerExecutorLossReasonInterface) 

	// public boolean checkSpeculatableTasks()
	CheckSpeculatableTasks() bool

	// public void increaseRunningTasks(int)
	IncreaseRunningTasks(a int) 

	// public void decreaseRunningTasks(int)
	DecreaseRunningTasks(a int) 
}

type SchedulerPool struct {
	JavaLangObject
}

// public java.lang.String logName()
func (jbobject *SchedulerPool) LogName() string {
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
func (jbobject *SchedulerPool) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String poolName()
func (jbobject *SchedulerPool) PoolName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "poolName", "java/lang/String")
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

// public int weight()
func (jbobject *SchedulerPool) Weight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "weight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int minShare()
func (jbobject *SchedulerPool) MinShare() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "minShare", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int runningTasks()
func (jbobject *SchedulerPool) RunningTasks() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "runningTasks", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int priority()
func (jbobject *SchedulerPool) Priority() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "priority", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int stageId()
func (jbobject *SchedulerPool) StageId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String name()
func (jbobject *SchedulerPool) Name() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "name", "java/lang/String")
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

// public org.apache.spark.scheduler.Pool parent()
func (jbobject *SchedulerPool) Parent() *SchedulerPool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "parent", "org/apache/spark/scheduler/Pool")
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

// public org.apache.spark.scheduler.SchedulingAlgorithm taskSetSchedulingAlgorithm()
func (jbobject *SchedulerPool) TaskSetSchedulingAlgorithm() *SchedulerSchedulingAlgorithm {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "taskSetSchedulingAlgorithm", "org/apache/spark/scheduler/SchedulingAlgorithm")
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
	unique_x := &SchedulerSchedulingAlgorithm{}
	unique_x.Callable = dst
	return unique_x
}

// public void addSchedulable(org.apache.spark.scheduler.Schedulable)
func (jbobject *SchedulerPool) AddSchedulable(a SchedulerSchedulableInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addSchedulable", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/Schedulable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeSchedulable(org.apache.spark.scheduler.Schedulable)
func (jbobject *SchedulerPool) RemoveSchedulable(a SchedulerSchedulableInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeSchedulable", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/Schedulable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.apache.spark.scheduler.Schedulable getSchedulableByName(java.lang.String)
func (jbobject *SchedulerPool) GetSchedulableByName(a string) *SchedulerSchedulable {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSchedulableByName", "org/apache/spark/scheduler/Schedulable", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &SchedulerSchedulable{}
	unique_x.Callable = dst
	return unique_x
}

// public void executorLost(java.lang.String, java.lang.String, org.apache.spark.scheduler.ExecutorLossReason)
func (jbobject *SchedulerPool) ExecutorLost(a string, b string, c SchedulerExecutorLossReasonInterface)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "executorLost", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("org/apache/spark/scheduler/ExecutorLossReason"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()

}

// public boolean checkSpeculatableTasks()
func (jbobject *SchedulerPool) CheckSpeculatableTasks() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "checkSpeculatableTasks", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void increaseRunningTasks(int)
func (jbobject *SchedulerPool) IncreaseRunningTasks(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "increaseRunningTasks", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void decreaseRunningTasks(int)
func (jbobject *SchedulerPool) DecreaseRunningTasks(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "decreaseRunningTasks", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


