package gospark

import "github.com/timob/javabind"

type SchedulerSchedulableInterface interface {

	// public abstract org.apache.spark.scheduler.Pool parent()
	Parent() *SchedulerPool

	// public abstract int weight()
	Weight() int

	// public abstract int minShare()
	MinShare() int

	// public abstract int runningTasks()
	RunningTasks() int

	// public abstract int priority()
	Priority() int

	// public abstract int stageId()
	StageId() int

	// public abstract java.lang.String name()
	Name() string

	// public abstract void addSchedulable(org.apache.spark.scheduler.Schedulable)
	AddSchedulable(a SchedulerSchedulableInterface) 

	// public abstract void removeSchedulable(org.apache.spark.scheduler.Schedulable)
	RemoveSchedulable(a SchedulerSchedulableInterface) 

	// public abstract org.apache.spark.scheduler.Schedulable getSchedulableByName(java.lang.String)
	GetSchedulableByName(a string) *SchedulerSchedulable

	// public abstract void executorLost(java.lang.String, java.lang.String, org.apache.spark.scheduler.ExecutorLossReason)
	ExecutorLost(a string, b string, c SchedulerExecutorLossReasonInterface) 

	// public abstract boolean checkSpeculatableTasks()
	CheckSpeculatableTasks() bool
}

type SchedulerSchedulable struct {
	JavaLangObject
}

// public abstract org.apache.spark.scheduler.Pool parent()
func (jbobject *SchedulerSchedulable) Parent() *SchedulerPool {
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

// public abstract int weight()
func (jbobject *SchedulerSchedulable) Weight() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "weight", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int minShare()
func (jbobject *SchedulerSchedulable) MinShare() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "minShare", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int runningTasks()
func (jbobject *SchedulerSchedulable) RunningTasks() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "runningTasks", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int priority()
func (jbobject *SchedulerSchedulable) Priority() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "priority", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int stageId()
func (jbobject *SchedulerSchedulable) StageId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract java.lang.String name()
func (jbobject *SchedulerSchedulable) Name() string {
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

// public abstract void addSchedulable(org.apache.spark.scheduler.Schedulable)
func (jbobject *SchedulerSchedulable) AddSchedulable(a SchedulerSchedulableInterface)  {
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

// public abstract void removeSchedulable(org.apache.spark.scheduler.Schedulable)
func (jbobject *SchedulerSchedulable) RemoveSchedulable(a SchedulerSchedulableInterface)  {
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

// public abstract org.apache.spark.scheduler.Schedulable getSchedulableByName(java.lang.String)
func (jbobject *SchedulerSchedulable) GetSchedulableByName(a string) *SchedulerSchedulable {
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

// public abstract void executorLost(java.lang.String, java.lang.String, org.apache.spark.scheduler.ExecutorLossReason)
func (jbobject *SchedulerSchedulable) ExecutorLost(a string, b string, c SchedulerExecutorLossReasonInterface)  {
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

// public abstract boolean checkSpeculatableTasks()
func (jbobject *SchedulerSchedulable) CheckSpeculatableTasks() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "checkSpeculatableTasks", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


